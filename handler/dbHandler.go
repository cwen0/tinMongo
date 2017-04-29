package handler

import (
	"net/http"

	gin "gopkg.in/gin-gonic/gin.v1"
)

func DbHome(c *gin.Context) {
	c.HTML(http.StatusOK, "db/home.html", nil)
}

func DbNewCollection(c *gin.Context) {
	c.HTML(http.StatusOK, "db/newCollection.html", nil)
}

func DbTransfer(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbTransfer.html", nil)
}

func DbExport(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbExport.html", nil)
}

func DbImport(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbImport.html", nil)
}

func DbUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbUsers.html", nil)
}
