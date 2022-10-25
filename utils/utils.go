package utils

import "golang.org/x/crypto/bcrypt"

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
