package user

import (
	"net/http"

	"gorm.io/gorm"
	"github.com/gin-gonic/gin"

	interfaces "banklineAPI/server/interfaces"
	mysql_utils "banklineAPI/server/db"
	utils "banklineAPI/server/utils"
)

func CreateAccount(r interfaces.Registry, db *gorm.DB) {
	user := interfaces.Credentials{Username: r.Username, Password: r.Password, Email: r.Email}
	_ = db.Create(&user)
}

func RegisterUser(c *gin.Context) {
	var input interfaces.Registry

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BAD REQUEST"})
		return
	}
	db := mysql_utils.ConnectDB(c)
	fd, _ := db.DB()
	defer fd.Close()

	// db.AutoMigrate(&interfaces.Credentials{})
	input.Password = utils.HidePassword([]byte(input.Password))
	CreateAccount(input, db)
}

func LoginUser(c *gin.Context) {
	var user interfaces.Login

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BAD REQUEST"})
		return
	}
	db := mysql_utils.ConnectDB(c)
	fd, _ := db.DB()
	defer fd.Close()

	key := ValidateUser(user.Username, user.Password, db)
	if key["message"] == "success" {
		token := utils.GenerateToken(&user)
		c.JSON(http.StatusOK, gin.H{"message": "success", "token": token})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": key["error"]})
	}
}

func AuthenticateUser(c *gin.Context) {
	var user interfaces.Authenticate

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "BAD REQUEST"})
		return
	}

	db := mysql_utils.ConnectDB(c)
	fd, _ := db.DB()
	defer fd.Close()

	key := ValidateUser(user.Username, user.Password, db)
	if key["message"] == "success" {
		if utils.ValidateToken(user.Token, user.Password) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "Token error"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"error": key["error"]})
	}
}
