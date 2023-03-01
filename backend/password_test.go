package main

import "testing"

func TestEncryption(t *testing.T) {
	const pw = "go-gators!"
	hash := encrypt(pw)

	if !comparePassword(pw, hash) {
		t.Errorf("Password mismatch")
	}
}
