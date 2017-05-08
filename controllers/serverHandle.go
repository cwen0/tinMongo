package controllers

import (
	"fmt"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/Sirupsen/logrus"
	"github.com/cwen0/tinMongo/models"
	"github.com/cwen0/tinMongo/utils"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	dbNames, err := mongo.DatabaseNames()
	if err != nil {
		logrus.Errorf("Get mongo database names failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	cmdLineOpts := bson.M{}
	if err = mongo.Run(bson.D{{"getCmdLineOpts", 1}}, &cmdLineOpts); err != nil {
		logrus.Errorf("Get mongo serverCmdLineOpts failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	buildInfo, err := mongo.BuildInfo()
	if err != nil {
		logrus.Errorf("Get mongo BuildInfo failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	connection := make(map[string]interface{})
	if host, ok := c.Get("host"); ok {
		fmt.Println(host)
		connection["Host"] = host
	}
	if port, ok := c.Get("port"); ok {
		connection["Port"] = port
	}
	h := utils.DefaultH(c)
	h["DBNames"] = dbNames
	h["ServerCmdLineOpts"] = cmdLineOpts
	h["BuildInfo"] = buildInfo
	h["Connection"] = connection
	h["GitHash"] = utils.TinMongoGitHash
	h["BuildTS"] = utils.TinMongoBuildTS

	c.HTML(200, "server/home", h)
}

func Status(c *gin.Context) {
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	ServerStatus := bson.M{}
	if err = mongo.Run(bson.D{{"serverStatus", 1}}, &ServerStatus); err != nil {
		logrus.Errorf("Get mongo serverStatus failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	h := utils.DefaultH(c)
	h["ServerStatus"] = ServerStatus
	c.HTML(200, "server/status", h)
}

func Databases(c *gin.Context) {
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	dbNames, err := mongo.DatabaseNames()
	if err != nil {
		logrus.Errorf("Get mongo database names failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	dbsStatus := make([]map[string]interface{}, 0)
	for _, name := range dbNames {
		status := make(map[string]interface{})
		db := mongo.DB(name)
		result := bson.M{}
		if err = db.Run(bson.D{{"dbstats", 1}}, &result); err != nil {
			logrus.Errorf("Get %d Status failed: %v", name, err)
			c.HTML(http.StatusInternalServerError, "errors/500", nil)
			return
		}
		status["Name"] = name
		status["DiskSize"] = "-"
		status["DataSize"] = "-"
		if v, ok := result["sizeOnDisk"]; ok {
			status["DiskSize"] = v
		}
		if v, ok := result["dataSize"]; ok {
			status["DataSize"] = v
		}
		status["StorageSize"] = result["storageSize"]
		status["IndexSize"] = result["indexSize"]
		status["Indexs"] = result["indexs"]
		status["Objects"] = result["objects"]
		status["Collections"] = result["collections"]
		dbsStatus = append(dbsStatus, status)
	}
	h := utils.DefaultH(c)
	h["DBsStatus"] = dbsStatus
	c.HTML(200, "server/databases", h)
}

func ProcessList(c *gin.Context) {
	c.HTML(200, "server/processList", map[string]interface{}{})
}

func Command(c *gin.Context) {
	c.HTML(200, "server/command", map[string]interface{}{})
}

func Execute(c *gin.Context) {
	c.HTML(200, "server/execute", map[string]interface{}{})
}

func Replication(c *gin.Context) {
	c.HTML(200, "server/replication", map[string]interface{}{})
}
