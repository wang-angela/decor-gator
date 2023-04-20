package controllers

// The code is formatted using this tutorial: https://blog.devgenius.io/sending-emails-with-golang-and-amazon-ses-31f25a0f2acb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"

	"github.com/decor-gator/backend/pkg/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func SendWelcomeEmail(destinationEmails []string) error {

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
		return err
	}

	return nil
}

func HelperForgotPassword(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	// Search for user by id; id=0 if user not found
	params := mux.Vars(r)["email"]
	filter := bson.D{{Key: "email", Value: params}}

	err := userColl.FindOne(context.TODO(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		// Throws error if no user exists
		msg := "User does not exist"

		err = json.NewEncoder(w).Encode(&msg)
		if err != nil {
			log.Fatalln("Error Encoding")
		}
		return
	} else if err != nil {
		// Throws error for other cases
		log.Fatal(err)
	}

	email := []string{user.Email}
	SendForgotPasswordEmail(email)

	err = json.NewEncoder(w).Encode("Email Sent")
	if err != nil {
		log.Fatalln("Error Encoding")
	}
}

func SendForgotPasswordEmail(destinationEmails []string) error {

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
		return err
	}

	return nil
}
