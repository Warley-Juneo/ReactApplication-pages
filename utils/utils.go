package utils

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	//gomail "gopkg.in/mail.v2"

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

func SendTokenEmail(user *interfaces.Login, token string) {
	// message := gomail.NewMessage()

	// message.SetHeader("from", user.Email)
	// message.SetHeader("to", user.Email)
	// message.SetHeader("Subject", "Messagem de envio de token")
	// message.SetBody("text", "Esse é um token privado, não passe para niguem: " + token)

	// a := gomail.NewDialer("smtp.gmail.com", 587, user.Email, "password")

	// if err := a.DialAndSend(message); err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }


	/////////////////// CODIGO NUMERO 2 ///////////////////
	fmt.Println("Nada a declarar")
}
