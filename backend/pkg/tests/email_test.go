package tests

import (
	"testing"

	"github.com/decor-gator/backend/pkg/controllers"
)

func TestSendWelcomeEmail(t *testing.T) {
	email := []string{"sgallic5@gmail.com"}

	err := controllers.SendWelcomeEmail(email)
	if err != nil {
		t.Error(err)
	}
}

func TestSendForgotPasswordEmail(t *testing.T) {
	email := []string{"sgallic5@gmail.com"}

	err := controllers.SendWelcomeEmail(email)
	if err != nil {
		t.Error(err)
	}
}
