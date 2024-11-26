package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"muyi2905/civic/backend/controllers"
)

// officialRoutes sets up the routes for ElectedOfficial CRUD operations
func officialRoutes(db *gorm.DB, r *gin.Engine) {
	r.POST("/officials", func(c *gin.Context) {
		controllers.CreateOfficial(db, c)
	})
	r.GET("/officials", func(c *gin.Context) {
		controllers.GetOfficials(db, c)
	})
	r.GET("/officials/:id", func(c *gin.Context) {
		controllers.GetOfficial(db, c)
	})
	r.PUT("/officials/:id", func(c *gin.Context) {
		controllers.UpdateOfficial(db, c)
	})
	r.DELETE("/officials/:id", func(c *gin.Context) {
		controllers.DeleteOfficial(db, c)
	})
}
