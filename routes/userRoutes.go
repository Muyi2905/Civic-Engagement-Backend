package routes

import (
	"muyi2905/civic/backend/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(db *gorm.DB, r *gin.Engine) {
	r.Group("api/users")
	{
		r.POST("/", controllers.Signup)
		r.POST("/", controllers.Login)
		r.GET("/", func(c *gin.Context) { controllers.GetUsers(db, c) })
		r.GET("/", func(c *gin.Context) { controllers.GetUserById(db, c) })
		r.PUT("/", func(c *gin.Context) { controllers.UpdateUserById(db, c) })
	}
}
