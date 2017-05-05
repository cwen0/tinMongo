package controllers

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/cwen0/tinMongo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login/login", nil)
}

func LoginPost(c *gin.Context) {
	session := sessions.Default(c)
	auth := &models.Auth{}
	var json = Wrapper{}
	if err := c.BindJSON(auth); err != nil {
		json.Errors = &Errors{Error{
			Status: http.StatusFound,
			Title:  "Please, fill out form correctly!!",
		}}
		c.JSON(http.StatusFound, json)
		return
	}
	mgoSess, err := auth.Connect()
	if err != nil {
		json.Errors = &Errors{Error{
			Status: http.StatusFound,
			Title:  "Authentication failed!!",
		}}
		logrus.Errorf("Login error, HostName: %s, Port: %d, UserName: %s, Password: %s, Database: %s", auth.HostName, auth.Port, auth.UserName, auth.Password, auth.Database)
		c.JSON(http.StatusFound, json)
		return
	}
	session.Set("mongo", mgoSess)

	session.Save()
	logrus.Info("Login sucess")
	c.JSON(http.StatusOK, json)
	return
}
