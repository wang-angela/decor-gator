package main

// Many aspects of these functions were taken from this tutorial: https://www.youtube.com/watch?v=hWmR8YtlFlE

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/decor-gator/backend/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

var loginKey = []byte("FakeKeyChangeLater")

var fakeUsers = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := fakeUsers[credentials.Username]

	if !ok || expectedPassword != credentials.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 2)

	aClaim := &Claims{
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, aClaim)
	tokenString, err := token.SignedString(loginKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
}

func Home(w http.ResponseWriter, r *http.Request) {

}

func Refresh(w http.ResponseWriter, r *http.Request) {

}
