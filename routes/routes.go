package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/okiprianto23/login-oauth/controllers"
	"github.com/okiprianto23/login-oauth/middlewares"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Auth routes
	r.POST("/auth/register", controllers.Register)
	r.POST("/auth/token", controllers.Login)
	r.POST("/auth/refresh", controllers.RefreshToken)

	// Protected routes
	protected := r.Group("/protected")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/data", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Protected Data"})
		})
	}

	return r
}
