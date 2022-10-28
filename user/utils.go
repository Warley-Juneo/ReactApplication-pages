package user

import (

	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"

	interfaces "banklineAPI/server/interfaces"
)

func ValidateUser(email string, password string, db *gorm.DB) map[string]interface{} {
	var user interfaces.Credentials

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return map[string]interface{}{"error": "User not found"}
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"error": "Wrong password"}
	}
	return map[string]interface{}{"message": "success"}
}

