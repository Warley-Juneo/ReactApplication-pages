package mysql_utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(c *gin.Context) *gorm.DB {
	dsn := "turboman_wjuneo-f:IHDmrH{vi}++@tcp(162.214.98.163:3306)/turboman_wjuneo-f?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "DB not connected"})
	}
	return db
}
