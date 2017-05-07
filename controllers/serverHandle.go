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
	//for v, _ := range ServerStatus {
	//fmt.Println(v)
	//}
	h := utils.DefaultH(c)
	h["ServerStatus"] = ServerStatus
	c.HTML(200, "server/status", h)
}

func Databases(c *gin.Context) {
	c.HTML(200, "server/databases", map[string]interface{}{})
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
