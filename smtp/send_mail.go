package smtp

import (
	"SampleApp/config"
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"text/template"
)

func SendMail(config *config.SMTPConfig, email string,templateFile string,templateInterface interface{}) {
	if config.ConfirmationUsername != "" {
		for _, registered := range config.Registered {
			if registered.Username == config.ConfirmationUsername {
				// Set up authentication information.
				from := fmt.Sprintf("%s@%s", registered.Username, config.Host)
				// Sender data.
				password := registered.Password

				// Receiver email address.
				to := []string{email}

				// smtp server configuration.
				smtpAddr := fmt.Sprintf("%s:%d", config.Host, config.Port)

				// Authentication.
				auth := smtp.PlainAuth("", from, password, config.Host)

				t, _ := template.ParseFiles(templateFile)

				var body bytes.Buffer

				mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
				body.Write([]byte(fmt.Sprintf("Subject: SampleApp - Confirm Email \n%s\n\n", mimeHeaders)))

				err := t.Execute(&body, templateInterface)
				if err != nil {
					log.Fatal("error parsing confirmation body ", err)
				}

				// Sending email
				err = smtp.SendMail(smtpAddr, auth, from, to, body.Bytes())
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("Email Sent!")
				break
			}
		}
	}
}
