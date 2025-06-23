package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oessona/job-tracker/internal/database"
	"github.com/oessona/job-tracker/internal/models"
	"net/http"
)

// CreateApplication - POST /applications
func CreateApplication(c *gin.Context) {
	var app models.Application

	if err := c.ShouldBindJSON(&app); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	app.UserID = userID.(uint)

	result := database.DB.Create(&app)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, app)
}

// GetApplications - GET /applications
func GetApplications(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}
	var apps []models.Application
	database.DB.Where("user_id = ?", userID).Find(&apps)

	c.JSON(http.StatusOK, apps)
}

// UpdateApplicationStatus - PATCH /applications/:id
func UpdateApplicationStatus(c *gin.Context) {
	id := c.Param("id")
	var app models.Application

	if err := database.DB.First(&app, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
		return
	}

	var input struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	app.Status = input.Status
	database.DB.Save(&app)

	c.JSON(http.StatusOK, app)
}

// GetApplicationStats - GET /applications/stats
func GetApplicationStats(c *gin.Context) {
	type Stat struct {
		Status string
		Count  int64
	}

	var stats []Stat
	database.DB.Model(&models.Application{}).
		Select("status, COUNT(*) as count").
		Group("status").
		Scan(&stats)

	c.JSON(http.StatusOK, stats)
}
