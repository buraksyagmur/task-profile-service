// task-profile-service/models/user.go

package models

import (
	"time"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username           string    `gorm:"unique;not null;type:varchar(255);index"`
	Email              string    `gorm:"unique;not null;type:varchar(255);index"`
	Password           string    `gorm:"not null;type:varchar(255)"`
	Verified           bool      `gorm:"default:false"`
	VerificationCode   string    `gorm:"type:varchar(6)"`
	VerificationExpiry time.Time `gorm:"type:datetime"`
}
