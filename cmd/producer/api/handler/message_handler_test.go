package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSendMessage(t *testing.T) {
	gin.SetMode(gin.TestMode)

	form := url.Values{}
	form.Add("message", "hello")
	router := gin.Default()
	router.POST("/insert", SendMessage)

	req, err := http.NewRequest(http.MethodPost, "/insert", strings.NewReader(form.Encode()))
	if err != nil {
		fmt.Printf("please enable kafka and zookeeper")
	}

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
