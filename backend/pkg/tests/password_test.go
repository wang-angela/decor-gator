package tests

import (
	"testing"

	"github.com/decor-gator/backend/pkg/utils"
)

func TestEncryption(t *testing.T) {
	const pw = "go-gators!"
	hash := utils.Encrypt(pw)

	if !utils.ComparePassword(pw, hash) {
		t.Errorf("Password mismatch")
	}
}
