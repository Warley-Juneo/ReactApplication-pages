package interfaces

import "gorm.io/gorm"

type Credentials struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

type Registry struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type Login struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Authenticate struct {
	Email string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
