package main

import (
	"net/http"
	"warung-pintar-test/cmd/socket/handler"
	"warung-pintar-test/cmd/socket/libs"

	"github.com/gin-gonic/gin"
)

func main() {
	go libs.H.Run()
	router := gin.Default()
	router.Static("/assets", "./public/assets")
	router.LoadHTMLGlob("public/*.html")

	router.GET("/ws", handler.ServerWs)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Run(":8080")
}
