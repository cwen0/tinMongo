package handler

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login/login.html", nil)
}
