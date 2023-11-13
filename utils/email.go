// task-profile-service/utils/email.go

package utils

import (
	"fmt"
	"log"
	"task-profile-service/config"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(to, verificationCode string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "burak.yagmur98@gmail.com")
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Email Verification for new email address")
	body := fmt.Sprintf("Your verification code is: %s", verificationCode)
	m.SetBody("text/plain", body)
	emailConfig := config.Cnfg.Email
	d := gomail.NewDialer(emailConfig.Host, emailConfig.Port, emailConfig.Username, emailConfig.Password)
	if err := d.DialAndSend(m); err != nil {
		log.Println("Error sending email:", err)
		return err
	}
	return nil
}
