package smtp

import (
	"SampleApp/config"
	"SampleApp/models"
)

func SendConfirmationMail(config *config.SMTPConfig, user *models.User) {
	SendMail(config, user.Email, "template/confirmation.html", user)
}
