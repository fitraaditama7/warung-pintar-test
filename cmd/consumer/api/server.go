package api

import (
	"warung-pintar-test/cmd/consumer/api/libs"
	"warung-pintar-test/cmd/consumer/api/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	libs.GetAllMessage()
	router := gin.Default()
	router.Use(cors.Default())
	message := router.Group("/message")
	{
		routes.MessageRouter(message)
	}

	router.Run(":4501")
}
