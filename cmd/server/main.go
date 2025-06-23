package main

import (
	"github.com/gin-gonic/gin"
	"github.com/oessona/job-tracker/internal/database"
	"github.com/oessona/job-tracker/internal/handlers"
	"github.com/oessona/job-tracker/internal/middleware"
	"github.com/oessona/job-tracker/internal/models"
)

func main() {
	r := gin.Default()

	database.Connect()
	database.DB.AutoMigrate(&models.User{}, &models.Application{})

	// ✅ Auth endpoints
	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)

	// ✅ Protected group
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.POST("/applications", handlers.CreateApplication)
	auth.GET("/applications", handlers.GetApplications)
	auth.PATCH("/applications/:id", handlers.UpdateApplicationStatus)

	r.GET("/applications/stats", handlers.GetApplicationStats) // можно оставить публичным

	r.Run(":8080")
}
