package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"rootkill/genie/database"
	"rootkill/genie/models"
	"rootkill/genie/routes"
)

func main() {
	// Initialize Database
	var db, err = database.InitDb()

	if err != nil {
		fmt.Println(err)
		return
	}

	// Run Database migrations for models
	db.AutoMigrate(&models.User{}, &models.Post{})

	// Initialize Gin router
	fmt.Println("Initialized gin router")
	router := gin.Default()

	// API routes
	fmt.Println("Attaching routes to the gin server")
	routes.SetupRoutes(router, db)

	// Start Gin Server
	fmt.Println("Server started at localhost:8080")
	router.Run(":8080")
}
