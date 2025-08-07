package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vigmiranda/coimobi-service/internal/property/controller"
	"github.com/vigmiranda/coimobi-service/pkg/configs"
	"net/http"
)

func main() {
	configs.ConnectDatabase()
	router := gin.Default()
	router.Use(corsMiddleware())

	colmobi := router.Group("/coimobi-admin")
	colmobi.GET("/property", controller.GetAllProperties)
	colmobi.GET("/property/:id", controller.GetProperty)
	colmobi.POST("/property", controller.InsertProperty)
	colmobi.PUT("/property/:id", controller.UpdateProperty)
	colmobi.DELETE("/property/:id", controller.DeleteProperty)

	router.GET("/health-check", func(c *gin.Context) { c.JSON(http.StatusOK, "OK") })
	router.Run(":8080")
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
