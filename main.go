package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var cnt int

func main() {
	cnt = 0
	router := gin.Default()
	router.LoadHTMLGlob("docs/*.html")
	router.GET("/", index)
	router.GET("/count", Count)
	router.Run()
}
func index(ctx *gin.Context) {
	cnt++
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}
func Count(ctx *gin.Context) {
	ctx.String(http.StatusOK, strconv.Itoa(cnt))
}
