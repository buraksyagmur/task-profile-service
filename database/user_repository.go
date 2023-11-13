// task-profile-service/database/user_repository.go
package database

import (
	"task-profile-service/models"
	"task-profile-service/utils"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func UpdateUserEmail(db *gorm.DB, userID uint, newEmail string) error {
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}
	verificationCode, err := utils.GenerateVerificationCode()
	if err != nil {
		return err
	}
	expirationTime := utils.GetExpirationTime()
	user.Email = newEmail
	user.Verified = false
	user.VerificationCode = verificationCode
	user.VerificationExpiry = expirationTime

	if err := db.Save(&user).Error; err != nil {
		return err
	}
	utils.SendVerificationEmail(newEmail, verificationCode)
	return nil
}
func VerifyUserPassword(db *gorm.DB, userID uint, password string) error {
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}
	return nil
}
func UpdateUserPassword(db *gorm.DB, userID uint, newPassword string) error {
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return err
	}

	if err := db.Model(&models.User{}).Where("ID = ?", userID).Update("Password", hashedPassword).Error; err != nil {
		return err
	}

	return nil
}
