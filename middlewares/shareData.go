package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

//SharedData fills in common data, such as user info, etc...
func SharedData() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if host := session.Get("host"); host != nil {
			c.Set("host", host)
		}

		if port := session.Get("port"); port != nil {
			c.Set("port", port)
		}
		c.Next()
	}
}
