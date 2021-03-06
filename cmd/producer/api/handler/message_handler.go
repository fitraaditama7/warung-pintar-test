package handler

import (
	"net/http"
	"warung-pintar-test/cmd/producer/api/libs"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
 * GET : '/message/send'
 *
 * @desc Send message to kafka
 *
 * @param  {string} message - Parameters for request
 *
 * @return {object} Request object
 */
func SendMessage(c *gin.Context) {
	message := c.Request.FormValue("message")
	var code int

	err := libs.SendMessage(message)
	if err != nil {
		logrus.Errorf("Unable to send message to kafka producer got error: %v", err)
		code = http.StatusInternalServerError
		message = "Gagal mengirim pesan"
	} else {
		code = http.StatusOK
		message = "Pesan berhasil dikirim"
	}

	result := gin.H{
		"code":    code,
		"message": message,
	}
	c.JSON(code, result)
	return
}
