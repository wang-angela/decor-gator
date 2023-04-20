// Code reference from: https://www.thepolyglotdeveloper.com/2017/03/authenticate-a-golang-api-with-json-web-tokens/
package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/decor-gator/backend/pkg/models"
	"github.com/decor-gator/backend/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/context"
	"github.com/mitchellh/mapstructure"
)

var PassPhrase = "SecretestSecret"

func CreateTokenEndpoint(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatalln("Error Decoding")
	}

	if !utils.JwtVerifyUserExists(user) {
		fmt.Print("User does not exist")
		return
	}

	if !utils.JwtVerifyPassword(user) {
		fmt.Print("Password incorrect")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Minute * 5).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(PassPhrase))
	if err != nil {
		log.Fatalln(err)
	}

	json.NewEncoder(w).Encode(models.JwtToken{Token: tokenStr})
}

func ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorizationHeader := r.Header.Get("authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte(PassPhrase), nil
				})
				if error != nil {
					json.NewEncoder(w).Encode(models.Exception{Message: error.Error()})
					return
				}
				if token.Valid {
					context.Set(r, "decoded", token.Claims)
					next(w, r)
				} else {
					json.NewEncoder(w).Encode(models.Exception{Message: "Invalid authorization token"})
				}
			}
		} else {
			json.NewEncoder(w).Encode(models.Exception{Message: "An authorization header is required"})
		}
	})
}

func VerifyEndpoint(w http.ResponseWriter, r *http.Request) {
	decoded := context.Get(r, "decoded")
	var user models.User
	mapstructure.Decode(decoded.(jwt.MapClaims), &user)
	json.NewEncoder(w).Encode(user)
}
