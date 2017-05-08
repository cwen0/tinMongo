package utils

import (
	"encoding/json"
	"html/template"
	"net/http"
	"reflect"

	"github.com/Sirupsen/logrus"
	"github.com/cwen0/tinMongo/models"
	humanize "github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

var TemplateFuncMap = template.FuncMap{
	"set":      Set,
	"equal":    Equal,
	"marshal":  Marshal,
	"humanize": Humanize,
}

func Set(args map[string]interface{}, key string, value interface{}) template.JS {
	args[key] = value
	return template.JS("")
}

func Equal(args ...interface{}) bool {
	return args[0] == args[1]
}

func DefaultH(c *gin.Context) gin.H {
	mongo, err := models.GetMongo()
	if err != nil {
		logrus.Errorf("Get mongo failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return nil
	}
	defer mongo.Close()
	dbNames, err := mongo.DatabaseNames()
	if err != nil {
		logrus.Errorf("Get mongo database names failed: %v", err)
		c.HTML(http.StatusInternalServerError, "errors/500", nil)
		return nil
	}
	host, _ := c.Get("host")
	port, _ := c.Get("port")
	return gin.H{
		"Host":    host,
		"Port":    port,
		"DBNames": dbNames,
	}
}

func Marshal(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}

func Humanize(v interface{}) interface{} {
	switch val := v.(type) {
	case uint32:
		return humanize.Bytes(uint64(val))
	case uint64:
		return humanize.Bytes(uint64(val))
	case int:
		return humanize.Bytes(uint64(val))
	case int32:
		return humanize.Bytes(uint64(val))
	case int64:
		return humanize.Bytes(uint64(val))
	case float64:
		return humanize.Bytes(uint64(val))
	case float32:
		return humanize.Bytes(uint64(val))
	case string:
		return val
	default:
		logrus.Errorf("Format Humanize failed, value:  %v, type: %v ", v, reflect.TypeOf(v))
		return v
	}
}
