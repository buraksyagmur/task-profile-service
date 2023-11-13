// task-profile-service/api/handlers/profile_handler.go
package handlers

import (
	"net/http"
	"task-profile-service/authentication"
	"task-profile-service/database"
	"task-profile-service/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get User Profile
// @Description Get the profile of the logged-in user.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.ResponseOK "User profile retrieved successfully"
// @Failure 401 {object} models.ResponseError "Unauthorized"
// @Router /profile [get]
func GetUserProfile(c *gin.Context) {
	authentication.AuthenticateJWT()(c)
	userID, _ := c.Get("user_id")
	username, _ := c.Get("username")
	useremail, _ := c.Get("useremail")

	c.JSON(http.StatusOK, gin.H{
		"userID":    userID,
		"username":  username,
		"useremail": useremail,
	})
}

// @Summary Update User Email
// @Description Update the email of the logged-in user.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param updateData body models.UpdateEmailData true "Update Email Data"
// @Success 200 {object} models.ResponseOK "Email updated successfully"
// @Failure 400 {object} models.ResponseError "Bad request"
// @Failure 401 {object} models.ResponseError "Unauthorized"
// @Router /profile/email [put]
func UpdateUserEmail(c *gin.Context) {
	authentication.AuthenticateJWT()(c)
	userID, _ := c.Get("user_id")
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ResponseError{Error: "Internal Server Error"})
		return
	}
	var updateEmailData models.UpdateEmailData
	if err := c.ShouldBindJSON(&updateEmailData); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
		return
	}
	if err := database.UpdateUserEmail(database.DB, userIDUint, updateEmailData.NewEmail); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{Error: "Failed to update email"})
		return
	}
	c.JSON(http.StatusOK, models.ResponseOK{Message: "Email updated successfully"})
}

// @Summary Update User Password
// @Description Update the password of the logged-in user.
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param updateData body models.UpdatePasswordData true "Update Password Data"
// @Success 200 {object} models.ResponseOK "Password updated successfully"
// @Failure 400 {object} models.ResponseError "Bad request"
// @Failure 401 {object} models.ResponseError "Unauthorized"
// @Router /profile/password [put]
func UpdateUserPassword(c *gin.Context) {
	authentication.AuthenticateJWT()(c)
	userID, _ := c.Get("user_id")
	userIDUint, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, models.ResponseError{Error: "Internal Server Error"})
		return
	}
	var updatePasswordData models.UpdatePasswordData
	if err := c.ShouldBindJSON(&updatePasswordData); err != nil {
		c.JSON(http.StatusBadRequest, models.ResponseError{Error: err.Error()})
		return
	}
	if updatePasswordData.NewPassword != updatePasswordData.ConfirmPassword {
		c.JSON(http.StatusBadRequest, models.ResponseError{Error: "New password and confirm password do not match"})
		return
	}
	if err := database.VerifyUserPassword(database.DB, userIDUint, updatePasswordData.OldPassword); err != nil {
		c.JSON(http.StatusUnauthorized, models.ResponseError{Error: "Incorrect old password"})
		return
	}
	if err := database.UpdateUserPassword(database.DB, userIDUint, updatePasswordData.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, models.ResponseError{Error: "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{Message: "Password updated successfully"})
}
