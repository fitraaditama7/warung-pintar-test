package routes

import (
	"warung-pintar-test/cmd/consumer/api/handler"

	"github.com/gin-gonic/gin"
)

func MessageRouter(router *gin.RouterGroup) {
	router.GET("/list", handler.GetAllMessage)
}
