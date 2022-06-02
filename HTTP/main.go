package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		req := c.Request
		loc := "https://" + req.Host + req.URL.Path
		if len(req.URL.RawQuery) > 0 {
			loc += "?" + req.URL.RawQuery
		}
		c.Redirect(http.StatusMovedPermanently, loc)
	})

	router.Run(":80")
}
