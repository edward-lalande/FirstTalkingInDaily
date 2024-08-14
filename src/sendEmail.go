package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendEmail(employe string) error {
	from := "test_go_mail@yahoo.com"
	password := "test-mail-g0"

	to := []string{strings.TrimSpace(employe)}

	smtpHost := "smtp.mail.yahoo.com"
	smtpPort := "587"

	message := []byte("Subject: Daily is in 10 minutes\r\n" +
		"\r\n" +
		"Hello, don't forget that you are the FIRST person talking today.\r\n" +
		"See you")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println("Failed to send email:", err)
		return fmt.Errorf("failed to send email: %w", err)
	}
	fmt.Println("Email Sent Successfully!")
	return nil
}
