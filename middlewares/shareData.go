package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//SharedData fills in common data, such as user info, etc...
func SharedData() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if mongo := session.Get("mongo"); mongo != nil {
			c.Set("mongo", mongo)
		}
		c.Set("mongo", "m")
		c.Next()
	}
}
