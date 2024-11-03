package email

import (
	"malai_agency/backend/env"
	"malai_agency/backend/services/logs"

	"gopkg.in/gomail.v2"
)

// SendEmail func used to send email
func SendEmail(mailContent EmailInput) error {
	m := gomail.NewMessage()
	m.SetHeader("From", env.SysEmail)
	m.SetHeader("To", mailContent.To...)
	m.SetHeader("Cc", mailContent.Cc...)
	m.SetHeader("Bcc", mailContent.Bcc...)
	m.SetHeader("Subject", mailContent.Subject)
	m.SetBody("text/html", mailContent.Message)
	MailCon := gomail.NewDialer("smtp.gmail.com", 465, env.SysEmail, env.SysEmailPassword)
	err := MailCon.DialAndSend(m)
	if err != nil {
		logs.Logs(" (SendEmail) send mail service Error:", err)
		return err
	}
	return nil
}
