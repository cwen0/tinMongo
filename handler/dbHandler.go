package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DbHome(c *gin.Context) {
	c.HTML(http.StatusOK, "db/home", map[string]interface{}{})
}

func DbNewCollection(c *gin.Context) {
	c.HTML(http.StatusOK, "db/newCollection", map[string]interface{}{})
}

func DbTransfer(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbTransfer", map[string]interface{}{})
}

func DbExport(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbExport", map[string]interface{}{})
}

func DbImport(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbImport", map[string]interface{}{})
}

func DbUsers(c *gin.Context) {
	c.HTML(http.StatusOK, "db/dbUsers", map[string]interface{}{})
}
