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

func Home(c *gin.Context) {
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	defer mongo.Close()
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
	if h == nil {
		return
	}
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
	defer mongo.Close()
	ServerStatus := bson.M{}
	if err = mongo.Run(bson.D{{"serverStatus", 1}}, &ServerStatus); err != nil {
		logrus.Errorf("Get mongo serverStatus failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	h := utils.DefaultH(c)
	if h == nil {
		return
	}
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
	defer mongo.Close()
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
		status["Indexes"] = result["indexes"]
		status["Objects"] = result["objects"]
		status["Collections"] = result["collections"]
		dbsStatus = append(dbsStatus, status)
	}
	h := utils.DefaultH(c)
	if h == nil {
		return
	}
	h["DBsStatus"] = dbsStatus
	c.HTML(200, "server/databases", h)
}

func CreateDatabase(c *gin.Context) {
	response := Wrapper{}
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Get mongo session failed: %v", err),
		}}
		c.JSON(http.StatusInsufficientStorage, response)
		return
	}
	dbName := strings.TrimSpace(c.PostForm("databaseName"))
	collectionName := strings.TrimSpace(c.PostForm("collectionName"))
	if dbName == "" || collectionName == "" {
		logrus.Error("CreateDatabase: form data is not correctly!")
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  "Please, fill out form correctly!",
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	defer mongo.Clone()
	db := mongo.DB(dbName)
	err = db.Run(bson.D{{"create", collectionName}}, nil)
	if err != nil {
		logrus.Error("Create database:collection [%s:%s] failed, DB: %v", dbName, collectionName, err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Create database:collection [%s:%s] failed, DB: %v", dbName, collectionName, err),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	logrus.Infof("Create database:collection [%s:%s] success", dbName, collectionName)
	c.JSON(http.StatusOK, response)
}

func DeleteDatabase(c *gin.Context) {
	response := Wrapper{}
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
	dbName := strings.TrimSpace(c.Param("dbName"))
	db := mongo.DB(dbName)
	err = db.DropDatabase()
	if err != nil {
		logrus.Error("Drop database [%s] failed: %v", dbName, err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Drop database [%s] failed: %v", dbName, err),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	logrus.Infof("Drop database [%s] success", dbName)
	c.JSON(http.StatusOK, response)
}

func ProcessList(c *gin.Context) {
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	defer mongo.Close()
	result := bson.M{}
	if err = mongo.Run(bson.D{{"currentOp", 1}, {"$all", 1}}, &result); err != nil {
		logrus.Errorf("Get ProcessList failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	h := utils.DefaultH(c)
	if h == nil {
		return
	}
	h["ProcessList"] = result["inprog"]
	c.HTML(200, "server/processList", h)
}

func Command(c *gin.Context) {
	h := utils.DefaultH(c)
	if h == nil {
		return
	}
	c.HTML(200, "server/command", h)
}

func ExecCommand(c *gin.Context) {
	response := Wrapper{}
	cmd := struct {
		Command string `json:"command"`
		DBName  string `json:"dbName"`
		// 	Format  string `json:"fotmat"`
	}{}
	if err := c.BindJSON(&cmd); err != nil {
		logrus.Errorf("ExecCommand bad required: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Please, fill out form corrently!!",
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if cmd.Command == "" || cmd.DBName == "" {
		logrus.Error("ExecCommand: form data is not corrrent")
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Please, fill out form corrently!!",
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var cmdBson interface{}
	if err := bson.UnmarshalJSON([]byte(cmd.Command), &cmdBson); err != nil {
		logrus.Errorf("UnmarshalJSON %s failed: %v", cmd.Command, err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  fmt.Sprintf("UnmarshalJSON %s failed: %v", cmd.Command, err),
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
	db := mongo.DB(cmd.DBName)
	result := bson.M{}
	if err := db.Run(cmdBson, &result); err != nil {
		logrus.Errorf("Run command [%v] failed: %v", cmd.Command, err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Run command [%v] failed: %v", cmd.Command, err),
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

func Execute(c *gin.Context) {
	h := utils.DefaultH(c)
	if h == nil {
		return
	}
	c.HTML(200, "server/execute", h)
}

func DoExecute(c *gin.Context) {
	response := Wrapper{}
	cmd := struct {
		Code   string `form:"code" json:"code"`
		DBName string `form:"dbName" json:"dbName"`
		Argus  string `from:"argus" json:"argus"`
	}{}
	if err := c.Bind(&cmd); err != nil {
		logrus.Errorf("DoExecute bad required: %v", err)
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
	defer mongo.Close()
	db := mongo.DB(cmd.DBName)
	result := bson.M{}
	if err := db.Run(bson.M{"eval": cmd.Code}, &result); err != nil {
		logrus.Errorf("Run JavaScript[%v] failed: %v", cmd.Code, err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Run JavaScript[%v] failed: %v", cmd.Code, err),
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

func Replication(c *gin.Context) {
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	defer mongo.Close()
	result := bson.M{}
	if err = mongo.Run(bson.M{"eval": `function () { return db.getReplicationInfo(); }`}, &result); err != nil {
		logrus.Errorf("Get getReplicationInfo failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	h := utils.DefaultH(c)
	h["ReplicationInfo"] = result
	c.HTML(200, "server/replication", h)
}
