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
		api.POST("/signup", func(c *gin.Context) { controllers.Signup(db, c) })
		api.POST("/login", func(c *gin.Context) { controllers.Login(db, c) })
		api.GET("/", func(c *gin.Context) { controllers.GetUsers(db, c) })
		api.GET("/:id", func(c *gin.Context) { controllers.GetUserById(db, c) }) // Added :id for user by ID
		api.PUT("/:id", func(c *gin.Context) { controllers.UpdateUser(db, c) })
	}
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/profile", func(c *gin.Context) { controllers.GetProfile(db, c) })

}
