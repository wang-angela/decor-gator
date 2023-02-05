package main

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func encrypt(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 15)
	if err != nil {
		log.Fatalln(err)
	}
	return string(bytes)
}

func comparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Fatal(err)
	}
	return true
}
