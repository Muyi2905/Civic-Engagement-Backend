package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"muyi2905/civic/backend/controllers"
	"muyi2905/civic/backend/middleware"
)

func UserRoutes(db *gorm.DB, r *gin.Engine) {
	r.Group("api/users")
	{
		r.POST("/", func(c *gin.Context) { controllers.Signup(db, c) })
		r.POST("/", func(c *gin.Context) { controllers.Login(db, c) })
		r.GET("/", func(c *gin.Context) { controllers.GetUsers(db, c) })
		r.GET("/", func(c *gin.Context) { controllers.GetUserById(db, c) })
		r.PUT("/", func(c *gin.Context) { controllers.UpdateUser(db, c) })
	}
	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/profile", func(c *gin.Context) { controllers.GetProfile(db, c) })

}
