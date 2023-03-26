package controllers

// The code is formatted using this tutorial: https://blog.devgenius.io/sending-emails-with-golang-and-amazon-ses-31f25a0f2acb

import (
	"fmt"
	"net/smtp"
)

func SendWelcomeEmail(destinationEmails []string) {

	// Creates necessary variables for the function.
	var (
		authUserName   = "AKIAWYOMFPS7EFQ4MFNL"
		authPassword   = "BGy07FXzzx3rQFXUUxzMvf/YKQsi97EtxzZyao70fDyb"
		smtpServerAddr = "email-smtp.us-east-1.amazonaws.com"
		smtpServerPort = "587"
		senderEmail    = "decorgators@gmail.com"
	)

	// Creates message for the email.
	msg := []byte("Subject: Welcome to DecorGators!\r\n" +
		"\r\n" +
		"You've successfully made an account with DecorGators!\r\n")

	// Gives authentification to send the email through AWS.
	auth := smtp.PlainAuth("", authUserName, authPassword, smtpServerAddr)

	// Sends email.
	err := smtp.SendMail(smtpServerAddr+":"+smtpServerPort,
		auth, senderEmail, destinationEmails, msg)

	// Catches error.
	if err != nil {
		fmt.Printf("Error to sending email: %s", err)
		return
	}
}

func SendForgotPasswordEmail(destinationEmails []string) {

	// Creates necessary variables for the function.
	var (
		authUserName   = "AKIAWYOMFPS7EFQ4MFNL"
		authPassword   = "BGy07FXzzx3rQFXUUxzMvf/YKQsi97EtxzZyao70fDyb"
		smtpServerAddr = "email-smtp.us-east-1.amazonaws.com"
		smtpServerPort = "587"
		senderEmail    = "decorgators@gmail.com"
	)

	// Creates message for the email.
	msg := []byte("Subject: Reset DecorGators Password\r\n" +
		"\r\n" +
		"Follow this link to reset your password:\r\n")

	// Gives authentification to send the email through AWS.
	auth := smtp.PlainAuth("", authUserName, authPassword, smtpServerAddr)

	// Sends email.
	err := smtp.SendMail(smtpServerAddr+":"+smtpServerPort,
		auth, senderEmail, destinationEmails, msg)

	// Catches error.
	if err != nil {
		fmt.Printf("Error to sending email: %s", err)
		return
	}
}
