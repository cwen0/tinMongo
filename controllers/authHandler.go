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
	var json = Wrapper{}
	if err := c.BindJSON(auth); err != nil {
		json.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Please, fill out form correctly!!",
		}}
		c.JSON(http.StatusBadRequest, json)
		return
	}
	mongo, err := auth.Connect()
	if err != nil {
		json.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Authentication failed!!",
		}}
		logrus.Errorf("Login error, HostName: %s, Port: %d, UserName: %s, Password: %s, Database: %s", auth.HostName, auth.Port, auth.UserName, auth.Password, auth.Database)
		c.JSON(http.StatusBadRequest, json)
		return
	}
	err = models.InitMongo(mongo)
	if err != nil {
		logrus.Errorf("Init mgo session failed: %v", err)
		json.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Init mgo session failed: %v", err),
		}}
		c.JSON(http.StatusInternalServerError, json)
	}
	//session.Set("mongo", fmt.Sprintf("%s:%s@%s:%d", auth.UserName, auth.Password, auth.UserName, auth.Port))
	session.Set("host", auth.HostName)
	session.Set("port", auth.Port)
	session.Save()
	logrus.Info("Login sucess")
	c.JSON(http.StatusOK, json)
}
