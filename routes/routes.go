package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"rootkill/genie/handlers"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize handlers
	userHandler := handlers.NewUserHandler(db)
	postHandler := handlers.NewPostHandler(db)

	// User Routes
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userHandler.GetUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// Post Routes
	postRoutes := router.Group("/posts")
	{
		postRoutes.GET("/", postHandler.GetPosts)
		postRoutes.POST("/", postHandler.CreatePost)
	}
}
