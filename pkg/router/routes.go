package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vigmiranda/coimobi-service/internal/property/controller"
	"net/http"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(corsMiddleware())

	// Health check route
	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	api := router.Group("/coimobi-admin")
	{
		// Property routes
		propertiesRoutes := api.Group("/property")
		{
			propertiesRoutes.GET("/all", controller.GetAllProperties)
			propertiesRoutes.GET("/:id", controller.GetProperty)
			propertiesRoutes.POST("/", controller.InsertProperty)
			propertiesRoutes.PUT("/:id", controller.UpdateProperty)
			propertiesRoutes.DELETE("/:id", controller.DeleteProperty)
		}
	}

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
