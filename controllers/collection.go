package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/Sirupsen/logrus"
	"github.com/cwen0/tinMongo/models"
	"github.com/cwen0/tinMongo/utils"
	"github.com/gin-gonic/gin"
)

func Document(c *gin.Context) {
	dbName := strings.TrimSpace(c.Param("dbName"))
	collection := strings.TrimSpace(c.Param("collection"))
	h := utils.DefaultH(c)
	if h == nil {
		return
	}
	h["DBName"] = dbName
	h["Collection"] = collection
	c.HTML(http.StatusOK, "collection/home", h)
}

func QueryDocument(c *gin.Context) {
	response := Wrapper{}
	query := struct {
		Query      string `json:"query"`
		DBName     string `json:"dbName"`
		Collection string `json:"collection"`
	}{}
	if err := c.BindJSON(&query); err != nil {
		logrus.Errorf("Query document failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Please, fill out form corrently!!",
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if query.Query == "" || query.DBName == "" || query.Collection == "" {
		logrus.Error("Query document: form data is not corrent")
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Please, fill out form corrently!!",
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var queryBson interface{}
	if err := bson.UnmarshalJSON([]byte(query.Query), &queryBson); err != nil {
		logrus.Errorf("UnamrshalJSON %s failed: %v", query.Query, err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  fmt.Sprintf("UnmarshalJSON %s failed: %v", query.Query, err),
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Get mongo session failed: %v", err),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	defer mongo.Close()
	result := []bson.M{}
	if err = mongo.DB(query.DBName).C(query.Collection).Find(queryBson).All(&result); err != nil {
		logrus.Errorf("Run Query[%v] failed: %s", queryBson, err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Run Query [%v] failed %v", queryBson, err),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	response.Datas = &Datas{Data{
		Type:    "json",
		Context: result,
	}}
	c.JSON(http.StatusOK, response)
}
