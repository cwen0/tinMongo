package controllers

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/Sirupsen/logrus"
	"github.com/cwen0/tinMongo/models"
	"github.com/cwen0/tinMongo/utils"
	"github.com/gin-gonic/gin"
)

type Record map[string]interface{}

func DBHome(c *gin.Context) {
	dbName := strings.TrimSpace(c.Param("dbName"))
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	defer mongo.Close()
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
		Size       string `json:"size"`
		FileCount  string `json:"fileCount"`
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
	defer mongo.Close()
	//capped := 0
	//if collectionInfo.IsCapped {
	//capped = 1
	//}
	db := mongo.DB(collectionInfo.DBName)
	size, _ := strconv.Atoi(collectionInfo.Size)
	max, _ := strconv.Atoi(collectionInfo.FileCount)
	cmdMap := bson.M{
		"create": collectionInfo.Collection,
		//	"capped": capped,
		"size": size,
		"max":  max,
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

func DBTransfer(c *gin.Context) {
	dbName := strings.TrimSpace(c.Param("dbName"))
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	defer mongo.Close()
	db := mongo.DB(dbName)
	collectionNames, err := db.CollectionNames()
	if err != nil {
		logrus.Errorf("Get collection names failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	h := utils.DefaultH(c)
	h["DBName"] = dbName
	h["CollectionNames"] = collectionNames
	c.HTML(http.StatusOK, "db/dbTransfer", h)
}

func ExecDBTransfer(c *gin.Context) {
	info := struct {
		Socket      string   `form:"socket" json:"socket"`
		Host        string   `form:"host" json:"host"`
		Port        string   `form:"port" json:"port"`
		IsAuth      bool     `from:"isAuth" json:"isAuth"`
		Username    string   `form:"username" json:"username"`
		Password    string   `form:"password" json:"password"`
		Collections []string `form:"collections" json:"collections"`
		DBName      string   `form:"dbName" json:"dbName"`
		IsCopyIndex bool     `form:"isCopyIndex" json:"isCopyIndex"`
	}{}
	response := Wrapper{}
	if err := c.Bind(&info); err != nil {
		logrus.Errorf("ExecDBTransfer bad requited: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Please. fill out form corrently!!",
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	url := info.Host + ":" + info.Port
	if info.Socket != "" {
		url = info.Socket
	}
	if info.IsAuth {
		url = "mongodb://" + info.Username + ":" + info.Password + "@" + url
	} else {
		url = "mongodb://" + url
	}
	session, err := mgo.Dial(url)
	if err != nil {
		logrus.Errorf("Get mongo  session from %v failed: %v", url, err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Get mongo session from %v failed: %v", url, err),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	server_url, ok := c.Get("url")
	if !ok {
		logrus.Error("Get server url failed")
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprint("Get server url failed"),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	for _, valuse := range info.Collections {
		cmd := bson.M{
			"cloneCollection": info.DBName + "." + valuse,
			"from":            server_url,
			"copyIndexes":     info.IsCopyIndex,
		}
		if err := session.Run(cmd, nil); err != nil {
			logrus.Errorf("Run copy collection[%v] failed: %v", cmd, err)
			response.Errors = &Errors{Error{
				Status: http.StatusInternalServerError,
				Title:  fmt.Sprintf("Run copy collection[%v] failed :%v", cmd, err),
			}}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	}
	c.JSON(http.StatusOK, response)
}

func DBExport(c *gin.Context) {
	dbName := strings.TrimSpace(c.Param("dbName"))
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	defer mongo.Close()
	db := mongo.DB(dbName)
	collectionNames, err := db.CollectionNames()
	if err != nil {
		logrus.Errorf("Get collection names failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	h := utils.DefaultH(c)
	h["DBName"] = dbName
	h["CollectionNames"] = collectionNames
	c.HTML(http.StatusOK, "db/dbExport", h)
}

func ExecDBExport(c *gin.Context) {
	info := struct {
		DBName      string `form:"dbName" json:"dbName"`
		Colls       string `form:"colls" json:"colls"`
		Collections []string
		IsDownload  bool `form:"isDownload" json:"isDownload"`
		IsGzip      bool `form:"isGzip" json:"isGzip"`
	}{}
	response := Wrapper{}
	if err := c.Bind(&info); err != nil {
		logrus.Errorf("ExecDBExport bad request: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Please. fill out form corrently!!",
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	colls := strings.Split(info.Colls, ",")
	for _, coll := range colls {
		info.Collections = append(info.Collections, strings.TrimSpace(coll))
	}
	var contexts string
	var countRows int
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	defer mongo.Close()
	db := mongo.DB(info.DBName)
	for _, collection := range info.Collections {
		coll := db.C(collection)
		indexs, _ := coll.Indexes()
		for _, index := range indexs {
			options := make(map[string]interface{})
			options["Unique"] = index.Unique
			optionsJson, err := json.Marshal(options)
			if err != nil {
				logrus.Errorf("Marshal [%v] failed: %v", options, err)
				c.HTML(http.StatusInternalServerError, "errors/500", nil)
				return
			}
			keys := make(map[string]interface{})
			keys["Key"] = index.Key
			keysJson, err := json.Marshal(keys)
			if err != nil {
				logrus.Errorf("Marshal [%v] failed: %v", keys, err)
				c.HTML(http.StatusInternalServerError, "errors/500", nil)
				return
			}
			contexts += fmt.Sprintf("\n/** {%s} indexs **/\n db.getCollection(\"%s\").ensureIndex(%s, %s);\n", collection, collection, keysJson, optionsJson)
		}
	}
	for _, collection := range info.Collections {
		contexts += fmt.Sprintf("\n/** %s  records **/\n", collection)
		iter := db.C(collection).Find(nil).Iter()
		defer iter.Close()
		var one Record
		for iter.Next(&one) {
			countRows++
			//	fmt.Println(one)
			//one, _ := one["data"].(Record)
			oneJson, err := json.Marshal(one)
			if err != nil {
				logrus.Errorf("Marshal [%v] failed: %v", one, err)
				c.HTML(http.StatusInternalServerError, "errors/500", nil)
				return
			}
			contexts += fmt.Sprintf("db.getCollection(\"%s\").insert(%s);\n", collection, oneJson)
		}
	}
	if info.IsDownload {
		var b bytes.Buffer
		w := gzip.NewWriter(&b)
		if info.IsGzip {
			w.Write([]byte(contexts))
			w.Flush()
			response.Datas = &Datas{Data{
				Type:    "gzip",
				Context: b,
			}}
		} else {
			response.Datas = &Datas{Data{
				Type:    "js",
				Context: contexts,
			}}
		}
	}
	c.JSON(http.StatusOK, response)
}

func DBImport(c *gin.Context) {
	h := utils.DefaultH(c)
	if h == nil {
		return
	}
	dbName := strings.TrimSpace(c.Param("dbName"))
	h["DBName"] = dbName
	c.HTML(http.StatusOK, "db/dbImport", h)
}

func ExecDBImport(c *gin.Context) {
	response := Wrapper{}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		logrus.Errorf("Get Import file failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  fmt.Sprintf("Get Import file failed: %v", err),
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Errorf("Read context from file failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  fmt.Sprintf("Read context from file failed: %v", err),
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	fileType := c.PostForm("fileType")
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
	db := mongo.DB(c.PostForm("dbName"))
	switch fileType {
	case "js":
		code := fmt.Sprintf("function (){ %s }", fileData)
		if err = db.Run(bson.M{"eval": code}, nil); err != nil {
			logrus.Errorf("Run Import js code failed: %v", err)
			response.Errors = &Errors{Error{
				Status: http.StatusInternalServerError,
				Title:  fmt.Sprintf("Run Import js code failed: %v", err),
			}}
			c.JSON(http.StatusInternalServerError, response)
			return
		}
	case "json":
		collection := c.PostForm("coll_name")
		coll := db.C(collection)
		lines := strings.Split(string(fileData), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line != "" {
				err = coll.Insert(line)
				if err != nil {
					logrus.Errorf("Run Import json data failed: %v", err)
					response.Errors = &Errors{Error{
						Status: http.StatusInternalServerError,
						Title:  fmt.Sprintf("Run Import json data failed: %v", err),
					}}
					c.JSON(http.StatusInternalServerError, response)
					return

				}
			}
		}
	default:
		logrus.Error("Import file type is not supported")
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Import file type is not supported",
		}}
		c.JSON(http.StatusBadRequest, response)
		return
	}
}

func DBUsers(c *gin.Context) {
	h := utils.DefaultH(c)
	if h == nil {
		return
	}
	dbName := strings.TrimSpace(c.Param("dbName"))
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	defer mongo.Clone()
	db := mongo.DB(dbName)
	err = db.AddUser("cwen1", "123", false)
	if err != nil {
		logrus.Error(err)
	}
	users := bson.M{}
	err = db.Run(bson.D{{"usersInfo", 1}}, &users)
	if err != nil {
		logrus.Errorf("Get users failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	h["DBName"] = dbName
	h["Users"] = users["users"]
	c.HTML(http.StatusOK, "db/dbUsers", h)
}

func DeleteDBUser(c *gin.Context) {
	response := Wrapper{}
	dbName := strings.TrimSpace(c.Param("dbName"))
	userName := strings.TrimSpace(c.Param("user"))
	if dbName == "" || userName == "" {
		logrus.Error("Delete user failed: request data is bad")
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  "Request data is bad !!",
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
	db := mongo.DB(dbName)
	err = db.RemoveUser(userName)
	if err != nil {
		logrus.Errorf("Delete user failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Delete user failed: %v", err),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func CreateDBUser(c *gin.Context) {
	response := Wrapper{}
	info := struct {
		DBName     string `form:"dbName" json:"dbName"`
		UserName   string `form:"user" json:"user"`
		Password   string `form:"password" json:"password"`
		IsReadonly bool   `form:"isReadonly" json:"isReadonly"`
	}{}
	if err := c.Bind(&info); err != nil {
		logrus.Errorf("Create User bad request: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  fmt.Sprintf("Please, fill out form corrently!!"),
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
	if err := mongo.DB(info.DBName).AddUser(info.UserName, info.Password, info.IsReadonly); err != nil {
		logrus.Errorf("Create User failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Create User failed: %v", err),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

func DBOperate(c *gin.Context) {
	dbName := strings.TrimSpace(c.Param("dbName"))

	h := utils.DefaultH(c)
	if h == nil {
		return
	}
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo session failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	defer mongo.Close()
	db := mongo.DB(dbName)
	collNames, err := db.CollectionNames()
	if err != nil {
		logrus.Errorf("Get collection names failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return
	}
	h["DBName"] = dbName
	h["Collections"] = collNames
	c.HTML(http.StatusOK, "db/dbOperate", h)
}

func DBClear(c *gin.Context) {
	response := Wrapper{}
	dbName := strings.TrimSpace(c.Param("dbName"))
	if dbName == "" {
		logrus.Error("DBClear DB name is empty")
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  fmt.Sprintf("Please, fill out form corrently!!"),
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
	db := mongo.DB(dbName)
	collNames, err := db.CollectionNames()
	if err != nil {
		logrus.Errorf("Get collection names failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Get collection names failed: %v", err),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	for _, coll := range collNames {
		db.C(coll).DropCollection()
	}
	c.JSON(http.StatusOK, response)
}

func DeleteCollection(c *gin.Context) {
	response := Wrapper{}
	dbName := strings.TrimSpace(c.Param("dbName"))
	if dbName == "" {
		logrus.Error("DBClear DB name is empty")
		response.Errors = &Errors{Error{
			Status: http.StatusBadRequest,
			Title:  fmt.Sprintf("Please, fill out form corrently!!"),
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
	db := mongo.DB(dbName)
	collName := strings.TrimSpace(c.Param("collName"))
	err = db.C(collName).DropCollection()
	if err != nil {
		logrus.Errorf("Drop collection  failed: %v", err)
		response.Errors = &Errors{Error{
			Status: http.StatusInternalServerError,
			Title:  fmt.Sprintf("Drop collection failed: %v", err),
		}}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
