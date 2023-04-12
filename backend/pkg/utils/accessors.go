// Accessors used by internal func without mux router
package utils

import (
	"github.com/decor-gator/backend/pkg/models"
)

// Used by CreatedTokenEndpoint() to check if users have an account
func JwtVerifyUserExists(user models.User) bool {
	var res models.User

	DB.Where("username = ?", user.Username).First(&res)
	if res.Username == "" {
		return false
	}

	return true
}

func JwtVerifyPassword(user models.User) bool {
	var res models.User

	DB.Where("username = ?", user.Username).First(&res)
	if !ComparePassword(user.Password, res.Password) {
		return false
	}

	return true
}
