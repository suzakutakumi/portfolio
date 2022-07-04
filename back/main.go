package main

import (
	"net/http"
	"portfolio/account"
	"portfolio/iot"
	"strconv"

	"github.com/gin-gonic/gin"
)

var cnt int

func main() {
	cnt = 0
	router := gin.Default()
	router.LoadHTMLGlob("docs/*.html")
	router.Static("/img", "docs/img")
	router.GET("/", index)
	router.GET("/count", Count)
	iot.Route(router)
	account.Route(router)
	//admin.Route(router)
	//admin.
	//router.Run("0.0.0.0:8080")
	router.RunTLS("0.0.0.0:8080", "/etc/letsencrypt/live/suzakutakumi.mydns.jp/fullchain.pem", "/etc/letsencrypt/live/suzakutakumi.mydns.jp/privkey.pem")
}
func index(ctx *gin.Context) {
	cnt++
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
func Count(ctx *gin.Context) {
	ctx.String(http.StatusOK, strconv.Itoa(cnt))
}
