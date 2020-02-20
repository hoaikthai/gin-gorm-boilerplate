package auth

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
