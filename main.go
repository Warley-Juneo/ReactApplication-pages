package main

import (
	"fmt"
	"net/http"
	"time"

	mysql_utils "banklineAPI/server/db"
	interfaces "banklineAPI/server/interfaces"
	utils "banklineAPI/server/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/register", RegisterUser)
	router.POST("/login", LoginUser)
	router.POST("/authenticate", AuthenticateUser)
	router.Run(":8080")

}

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

	key := Validate(user.Username, user.Password, db)
	if key["message"] == "success" {
		token := GenerateToken(&user)
		c.JSON(http.StatusOK, gin.H{"message": "success", "token": token})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": key["error"]})
	}
}

func Validate(username string, password string, db *gorm.DB) map[string]interface{} {
	var user interfaces.Credentials

	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return map[string]interface{}{"error": "User not found"}
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return map[string]interface{}{"error": "Wrong password"}
	}
	return map[string]interface{}{"message": "success"}
}

func GenerateToken(user *interfaces.Login) string {
	tokenContent := jwt.MapClaims{
		"password": user.Password,
		"expire":   time.Now().Add(time.Minute ^ 1).Unix(),
	}
	fmt.Println(tokenContent["expire"])
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	utils.HandleError(err)
	return token
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

	key := Validate(user.Username, user.Password, db)
	if key["message"] == "success" {
		if ValidateToken(user.Token, user.Password) {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"error": "Token error"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"error": key["error"]})
	}
}

func ValidateToken(token string, password string) bool {
	tokenData := jwt.MapClaims{}
	jwt.ParseWithClaims(token, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})
	if tokenData["password"] == password {
		return true
	} else {
		return false
	}
}
