package main

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupTestRouter(db *gorm.DB) *gin.Engine {
	// Create Gin Test Server Configuration
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	routes.SetUpRoutes(router, db)
	return router
}
