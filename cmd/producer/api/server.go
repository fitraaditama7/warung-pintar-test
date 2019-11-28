package api

import (
	"net/http"
	"warung-pintar-test/cmd/producer/api/routes"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.Static("/assets", "./public/assets")
	router.LoadHTMLGlob("public/*.html")

	message := router.Group("/message")
	{
		routes.MessageRouter(message)
	}
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.Run(":4500")
}
