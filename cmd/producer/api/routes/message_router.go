package routes

import (
	"warung-pintar-test/cmd/producer/api/handler"

	"github.com/gin-gonic/gin"
)

func MessageRouter(router *gin.RouterGroup) {
	router.POST("/send", handler.SendMessage)
}
