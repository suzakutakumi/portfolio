package iot

import (
	"fmt"
	"net/http"
	"os"
	"portfolio/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type IoTData struct {
	Id     int    `json:"id" db:"id" binding:"required"`
	Name   string `json:"name" db:"nickName" binding:"required"`
	Kind   int    `json:"kind db:"kind" binding:"required"`
	UserId string `json:"userId" db:"userId"`
}

func Route(router *gin.Engine) {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf(".envを読み込めませんでした: %v", err)
		return
	}

	id := os.Getenv("adminId")
	passwd := os.Getenv("adminPass")
	iRouter := router.Group("/iot", gin.BasicAuth(gin.Accounts{
		id: passwd,
	}))

	iRouter.GET("/", index)
	iRouter.GET("/data", getIoT)
	iRouter.POST("/data", register)
}
func index(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	c.JSON(200, gin.H{"message": "Hello " + user})
}
func getIoT(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	var iot []IoTData
	if err := db.Select(&iot, "select * from iot where userId=?", user); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, iot)
}

func register(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	var iot IoTData
	if err := c.ShouldBindJSON(&iot); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	iot.UserId = user
	err := db.Push("insert into iot values(?,?,?,?,?)", iot.Id, iot.Name, iot.Kind, iot.UserId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
