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
		authorized.GET("/processList", controllers.ProcessList)
		authorized.GET("/command", controllers.Command)
		authorized.GET("/execute", controllers.Execute)
		authorized.GET("/replication", controllers.Replication)

		authorized.GET("/db/home", controllers.DbHome)
		authorized.GET("/db/newCollection", controllers.DbNewCollection)
		authorized.GET("/db/dbTransfer", controllers.DbTransfer)
		authorized.GET("/db/dbExport", controllers.DbExport)
		authorized.GET("/db/dbImport", controllers.DbImport)
		authorized.GET("/db/dbUsers", controllers.DbUsers)
	}
	return r
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if mongo, _ := c.Get("mongo"); mongo != nil {
			c.Next()
		} else {
			logrus.Warnf("User not authorized to visit %s", c.Request.RequestURI)
			c.HTML(http.StatusForbidden, "errors/403", nil)
			c.Abort()
		}
	}
}
