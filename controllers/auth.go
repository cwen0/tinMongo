package controllers

import (
	"fmt"
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
	response := Wrapper{}
	if err := c.BindJSON(auth); err != nil {
		logrus.Errorf("BindJSON failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Please, fill out form correctly!!",
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	mongo, url, err := auth.Connect()
	if err != nil {
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Authentication failed!!",
		}}
		logrus.Errorf("Login error, HostName: %s, Port: %d, UserName: %s, Password: %s, Database: %s", auth.HostName, auth.Port, auth.UserName, auth.Password, auth.Database)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = models.InitMongo(mongo)
	if err != nil {
		logrus.Errorf("Init mgo session failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Init mgo session failed: %v", err),
		}}
		c.JSON(http.StatusInternalServerError, response)
	}
	//session.Set("mongo", fmt.Sprintf("%s:%s@%s:%d", auth.UserName, auth.Password, auth.UserName, auth.Port))
	session.Set("host", auth.HostName)
	session.Set("port", auth.Port)
	session.Set("url", url)
	session.Save()
	logrus.Info("Login sucess")
	c.JSON(http.StatusOK, response)
}
