package utils

import (
	"encoding/json"
	"html/template"

	"github.com/gin-gonic/gin"
)

var TemplateFuncMap = template.FuncMap{
	"set":     Set,
	"equal":   Equal,
	"marshal": Marshal,
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
