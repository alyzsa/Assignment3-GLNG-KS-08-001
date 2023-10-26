package handlers

import (
	"net/http"

	"github.com/alyzsa/Assignment3-GLNG-KS-08-001/database"
	"github.com/alyzsa/Assignment3-GLNG-KS-08-001/models"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/getweather", RetrieveWeatherData)
	return router
}

func RetrieveWeatherData(c *gin.Context) {
	if database.DB == nil {
		if err := database.InitializeDatabase(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize the database: " + err.Error()})
			return
		}
		defer database.CloseDatabase()
	}

	latestWeatherData, err := models.GetLatestWeatherData(database.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch weather data: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, latestWeatherData)
}
