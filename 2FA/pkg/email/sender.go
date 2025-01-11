package email

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

func SendVerificationCode(code string) {
	senderEmail := os.Getenv("SENDER_EMAIL_ADDRESS")
	senderPassword := os.Getenv("SENDER_EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_SERVER")
	smtpPort := 587
	recipientEmail := os.Getenv("RECIPIENT_EMAIL_ADDRESS")

	m := gomail.NewMessage()
	m.SetHeader("From", senderEmail)
	m.SetHeader("To", recipientEmail)
	m.SetHeader("Subject", "Your Verification Code")
	m.SetBody("text/plain", fmt.Sprintf("Your verification code is: %s", code))

	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)
	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error sending verification code: %v", err)
	}
}
