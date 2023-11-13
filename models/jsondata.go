package models
type UpdateEmailData struct {
	NewEmail string `json:"new_email" binding:"required,email"`
}

type UpdatePasswordData struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
type Config struct {
	Database DatabaseConfig `json:"database"`
	Email    EmailConfig    `json:"email"`
}
type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Name     string `json:"name"`
}
type EmailConfig struct {
	Host     string `json:"smtp_server"`
	Port     int    `json:"smtp_port"`
	Username string `json:"username"`
	Password string `json:"password"`
}