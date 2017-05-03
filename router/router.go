package router

import (
	"github.com/cwen0/tinMongo/controllers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/login", controllers.Login)
	r.GET("/home", controllers.Home)
	r.GET("/status", controllers.Status)
	r.GET("/databases", controllers.Databases)
	r.GET("/processList", controllers.ProcessList)
	r.GET("/command", controllers.Command)
	r.GET("/execute", controllers.Execute)
	r.GET("/replication", controllers.Replication)

	r.GET("/db/home", controllers.DbHome)
	r.GET("/db/newCollection", controllers.DbNewCollection)
	r.GET("/db/dbTransfer", controllers.DbTransfer)
	r.GET("/db/dbExport", controllers.DbExport)
	r.GET("/db/dbImport", controllers.DbImport)
	r.GET("/db/dbUsers", controllers.DbUsers)
	return r
}
