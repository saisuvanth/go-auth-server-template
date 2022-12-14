package utils

import (
	"log"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(ch chan string, password string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		log.Fatal(err)
	}
	ch <- string(hash)
}

func ComparePassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func GenToken(ch chan string, userId string) {
	token := jwt.EncodeSegment([]byte(userId))
	ch <- token
}
