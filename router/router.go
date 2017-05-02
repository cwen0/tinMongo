package router

import (
	"github.com/cwen0/tinMongo/handler"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/login", handler.Login)
	r.GET("/home", handler.Home)
	r.GET("/status", handler.Status)
	r.GET("/databases", handler.Databases)
	r.GET("/processList", handler.ProcessList)
	r.GET("/command", handler.Command)
	r.GET("/execute", handler.Execute)
	r.GET("/replication", handler.Replication)

	r.GET("/db/home", handler.DbHome)
	r.GET("/db/newCollection", handler.DbNewCollection)
	r.GET("/db/dbTransfer", handler.DbTransfer)
	r.GET("/db/dbExport", handler.DbExport)
	r.GET("/db/dbImport", handler.DbImport)
	r.GET("/db/dbUsers", handler.DbUsers)
	return r
}
