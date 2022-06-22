package admin

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Route(router *gin.Engine) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf(".envを読み込めませんでした: %v", err)
		return
	}
	id := os.Getenv("id")
	passwd := os.Getenv("passwd")
	authorized := router.Group("/admin", gin.BasicAuth(gin.Accounts{
		id: passwd,
	}))
	authorized.GET("/", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		c.JSON(200, gin.H{"message": "Hello " + user})
	})
}
