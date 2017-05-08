package utils

import (
	"encoding/json"
	"html/template"
	"reflect"

	"github.com/Sirupsen/logrus"
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
	host, _ := c.Get("host")
	port, _ := c.Get("port")
	return gin.H{
		"Host": host,
		"Port": port,
	}
}

func Marshal(v interface{}) template.JS {
	a, _ := json.Marshal(v)
	return template.JS(a)
}

func Humanize(v interface{}) interface{} {
	//val, ok := v.(uint64)
	//if !ok {
	//logrus.Errorf("Format Humanize failed: %v", v)
	//return v
	//}
	//	return humanize.Bytes(val)
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
