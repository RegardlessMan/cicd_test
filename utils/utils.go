package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(pwd string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	return string(password), err
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})
	signedString, err := token.SignedString([]byte("secret"))
	return "Bearer " + signedString, err
}

func CheckPassword(inputPwd string, dbPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPwd), []byte(inputPwd))
	return err == nil
}
