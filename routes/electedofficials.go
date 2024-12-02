package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"muyi2905/civic/backend/controllers"
)

func OfficialRoutes(db *gorm.DB, r *gin.Engine) {
	api := r.Group("/api")
	{
		officials := api.Group("/officials")
		{
			officials.POST("", func(c *gin.Context) {
				controllers.CreateOfficial(db, c)
			})
			officials.GET("", func(c *gin.Context) {
				controllers.GetOfficials(db, c)
			})
			officials.GET("/:id", func(c *gin.Context) {
				controllers.GetOfficial(db, c)
			})
			officials.PUT("/:id", func(c *gin.Context) {
				controllers.UpdateOfficial(db, c)
			})
			officials.DELETE("/:id", func(c *gin.Context) {
				controllers.DeleteOfficial(db, c)
			})
		}
	}
}
