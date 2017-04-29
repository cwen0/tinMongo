package handler

import (
	gin "gopkg.in/gin-gonic/gin.v1"
)

func Home(c *gin.Context) {
	c.HTML(200, "server/home.html", nil)
}

func Status(c *gin.Context) {
	c.HTML(200, "server/status.html", nil)
}

func Databases(c *gin.Context) {
	c.HTML(200, "server/databases.html", nil)
}

func ProcessList(c *gin.Context) {
	c.HTML(200, "server/processList.html", nil)
}

func Command(c *gin.Context) {
	c.HTML(200, "server/command.html", nil)
}

func Execute(c *gin.Context) {
	c.HTML(200, "server/execute.html", nil)
}

func Replication(c *gin.Context) {
	c.HTML(200, "server/replication.html", nil)
}
