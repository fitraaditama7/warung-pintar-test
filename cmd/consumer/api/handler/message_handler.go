package handler

import (
	"net/http"
	"warung-pintar-test/cmd/consumer/api/libs"
	"warung-pintar-test/cmd/consumer/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*
 * GET : '/message/list'
 *
 * @desc Get list message in kafka
 *
 *
 * @return {object} Request object
 */
func GetAllMessage(c *gin.Context) {
	var code int
	var data interface{}

	err := libs.GetAllMessage()
	if err != nil {
		logrus.Errorf("Unable to get all message from kafka got error: %v", err)
		code = http.StatusInternalServerError
		data = nil
	} else {
		code = http.StatusOK
		data = &config.LISTMESSAGE
	}
	result := gin.H{
		"code": code,
		"data": data,
	}

	c.JSON(code, result)
	return
}
