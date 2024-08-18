package main

import (
	"fmt"
	"net/smtp"
	"strings"
)

func SendEmail(email string, password string, employe string) error {
	to := []string{strings.TrimSpace(employe)}

	smtpHost := "smtp.mail.yahoo.com"
	smtpPort := "587"

	println("email: [", email, "]", "\npassword: [", password, "]", "\nemploye: [", employe, "]\n")

	message := []byte("Subject: Daily is in 10 minutes\r\n" +
		"\r\n" +
		"Hello, don't forget that you are the FIRST person talking today.\r\n" +
		"See you")

	auth := smtp.PlainAuth("", email, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, email, to, message)
	if err != nil {
		fmt.Println("Failed to send email:", err)
		return fmt.Errorf("failed to send email: %w", err)
	}
	fmt.Println("Email Sent Successfully!")
	return nil
}
