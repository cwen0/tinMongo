package router

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/cwen0/tinMongo/controllers"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/login", controllers.LoginGet)
	r.POST("/login", controllers.LoginPost)
	authorized := r.Group("/server")
	authorized.Use(AuthRequired())
	{
		authorized.GET("/home", controllers.Home)
		authorized.GET("/status", controllers.Status)
		authorized.GET("/databases", controllers.Databases)
		authorized.POST("/newDatabase", controllers.CreateDatabase)
		authorized.POST("/database/:dbName/delete", controllers.DeleteDatabase)
		authorized.GET("/processList", controllers.ProcessList)
		authorized.GET("/command", controllers.Command)
		authorized.POST("/command", controllers.ExecCommand)
		authorized.GET("/execute", controllers.Execute)
		authorized.POST("/execute", controllers.DoExecute)
		authorized.GET("/replication", controllers.Replication)

		authorized.GET("/db/home/:dbName", controllers.DBHome)
		authorized.GET("/db/newCollection/:dbName", controllers.DBNewCollection)
		authorized.POST("/db/newCollection", controllers.ExecDBNewCollection)
		authorized.GET("/db/dbTransfer/:dbName", controllers.DBTransfer)
		authorized.POST("/db/dbTransfer", controllers.ExecDBTransfer)
		authorized.GET("/db/dbExport/:dbName", controllers.DBExport)
		authorized.POST("/db/dbExport", controllers.ExecDBExport)
		authorized.GET("/db/dbImport/:dbName", controllers.DBImport)
		authorized.POST("/db/dbImport", controllers.ExecDBImport)
		authorized.GET("/db/dbUsers/:dbName", controllers.DBUsers)
		authorized.POST("/db/dbUsers/:dbName/:user/delete", controllers.DeleteDBUser)
		authorized.POST("/db/newDBUser", controllers.CreateDBUser)
		authorized.GET("/db/dbOperate/:dbName", controllers.DBOperate)
		authorized.POST("/db/dbOperate/:dbName/clear", controllers.DBClear)
		authorized.POST("/db/collection/:dbName/:collName/delete", controllers.DeleteCollection)
		authorized.GET("/collection/home/:dbName/:collection", controllers.Document)
		authorized.POST("/collection/document/query", controllers.QueryDocument)
		authorized.POST("/collection/document/delete", controllers.DeleteDocument)
		authorized.POST("/collection/document/update", controllers.UpdateDocument)
		authorized.GET("/collection/document/insert/:dbName/:collection", controllers.InsertDocument)
		authorized.POST("/collection/document/insert", controllers.ExecInsertDocument)
	}
	return r
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Get("host"); ok {
			c.Next()
		} else {
			logrus.Warnf("User not authorized to visit %s", c.Request.RequestURI)
			c.HTML(http.StatusForbidden, "errors/403", nil)
			c.Abort()
		}
	}
}
