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

func DBHome(c *gin.Context) {
	dbName := strings.TrimSpace(c.Param("dbName"))
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	db := mongo.DB(dbName)
	result := bson.M{}
	if err = db.Run(bson.D{{"dbStats", 1}}, &result); err != nil {
		logrus.Errorf("Get database [%s] status failed: %v", dbName, err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	h := utils.DefaultH(c)
	if h == nil {
		return
	}
	result["indexSize"] = utils.Humanize(result["indexSize"])
	result["storageSize"] = utils.Humanize(result["storageSize"])
	result["dataSize"] = utils.Humanize(result["dataSize"])

	h["DBStats"] = result
	h["DBName"] = dbName
	c.HTML(http.StatusOK, "db/home", h)
}

func DBNewCollection(c *gin.Context) {
	dbName := strings.TrimSpace(c.Param("dbName"))
	h := utils.DefaultH(c)
	if h == nil {
		return
	}
	h["DBName"] = dbName
	c.HTML(http.StatusOK, "db/newCollection", h)
}

func ExecDBNewCollection(c *gin.Context) {
	response := Wrapper{}
	collectionInfo := struct {
		DBName     string `json:"dbName"`
		Collection string `json:"collection"`
		IsCapped   bool   `json:"isCapped"`
		Size       int    `json:"size"`
		FileCount  int    `json:"fileCount"`
	}{}
	if err := c.BindJSON(&collectionInfo); err != nil {
		logrus.Errorf("BindJSON failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Please, fill out form corrently!!",
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
	//capped := 0
	//if collectionInfo.IsCapped {
	//capped = 1
	//}
	db := mongo.DB(collectionInfo.DBName)
	cmdMap := bson.M{
		"create": collectionInfo.Collection,
		//	"capped": capped,
		"size": collectionInfo.Size,
		"max":  collectionInfo.FileCount,
	}

	if err = db.Run(cmdMap, nil); err != nil {
		logrus.Errorf("Run create collection[%v] failed: %v", cmdMap, err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Run create collection[%v] failed: %v", cmdMap, err),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func DbTransfer(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbTransfer", map[string]interface{}{})
}

func DbExport(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbExport", map[string]interface{}{})
}

func DbImport(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbImport", map[string]interface{}{})
}

func DbUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbUsers", map[string]interface{}{})
}
