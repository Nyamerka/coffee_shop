package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", RegisterHandler)
		authGroup.POST("/login", LoginHandler)
	}

	protectedGroup := router.Group("/protected")
	protectedGroup.Use(JWTAuthMiddleware())
	{
		protectedGroup.GET("/resource", CasbinMiddleware("resource_name", "read"), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Access granted"})
		})
	}
}
