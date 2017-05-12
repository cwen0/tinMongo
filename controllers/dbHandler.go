package controllers

import (
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

func DbNewCollection(c *gin.Context) {
	c.HTML(http.StatusOK, "db/newCollection", map[string]interface{}{})
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
