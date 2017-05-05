package controllers

import (
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(200, "server/home", map[string]interface{}{})
}

func Status(c *gin.Context) {
	c.HTML(200, "server/status", map[string]interface{}{})
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
