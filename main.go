package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vigmiranda/coimobi-service/api/configs"
	"github.com/vigmiranda/coimobi-service/api/controllers"
	"net/http"
)

func main() {
	configs.ConnectDatabase()
	router := gin.Default()

	colmobi := router.Group("/coimobi-admin")
	colmobi.GET("/property", controllers.GetAllProperties)
	colmobi.POST("/property", controllers.InsertProperty)
	colmobi.PUT("/property/:id", controllers.UpdateProperty)
	colmobi.DELETE("/property/:id", controllers.DeleteProperty)

	router.GET("/health-check", func(c *gin.Context) { c.JSON(http.StatusOK, "OK") })
	router.Run(":8080")
}
