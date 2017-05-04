package controllers

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/cwen0/tinMongo/models"
	"github.com/gin-gonic/gin"
)

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login/login", nil)
}

func LoginPost(c *gin.Context) {
	//	_ = sessions.Default(c)
	authConfig := &models.AuthConfig{}
	if err := c.BindJSON(authConfig); err != nil {
		logrus.Error(err)
	}
	logrus.Info(authConfig)
}
