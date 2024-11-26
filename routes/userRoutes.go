package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"muyi2905/civic/backend/controllers"
	"muyi2905/civic/backend/middleware"
)

func UserRoutes(db *gorm.DB, r *gin.Engine) {

	api := r.Group("/api/users")
	{
		// Public routes
		api.POST("/signup", func(c *gin.Context) { controllers.Signup(db, c) }) // Register new users
		api.POST("/login", func(c *gin.Context) { controllers.Login(db, c) })   // User login

		// User information routes
		api.GET("/", func(c *gin.Context) { controllers.GetUsers(db, c) })       // List all users (consider protecting)
		api.GET("/:id", func(c *gin.Context) { controllers.GetUserById(db, c) }) // Get user by ID
		api.PUT("/:id", func(c *gin.Context) { controllers.UpdateUser(db, c) })  // Update user by ID
	}

	// Authenticated routes
	authorized := api.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/profile", func(c *gin.Context) { controllers.GetProfile(db, c) }) // Get user profile
}
