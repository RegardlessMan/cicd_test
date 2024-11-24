package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "Login Success!",
			})
		})

		auth.POST("/register", func(c *gin.Context) {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "Register Success!",
			})
		})
	}

	return r
}
