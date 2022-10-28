package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	user "banklineAPI/server/user"
)

func StartApi () {

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/register", user.RegisterUser)
	router.POST("/login", user.LoginUser)
	router.POST("/authenticate", user.AuthenticateUser)
	router.Run(":8080")
}
