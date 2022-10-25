package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	interfaces "banklineAPI/server/interfaces"
)

func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HidePassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	HandleError(err)
	return string(hash)
}

func GenerateToken(user *interfaces.Login) string {
	tokenContent := jwt.MapClaims{
		"password": user.Password,
		"expire":   time.Now().Add(time.Minute ^ 1).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	HandleError(err)
	return token
}

func ValidateToken(token string, password string) bool {
	tokenData := jwt.MapClaims{}
	jwt.ParseWithClaims(token, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})
	if tokenData["password"] == password {
		if tokenData["expire"].(float64) > float64(time.Now().Local().Unix()) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}
