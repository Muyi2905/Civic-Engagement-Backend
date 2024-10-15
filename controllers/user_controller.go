package controllers

import (
	"muyi2905/civic/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

func GetUsers(db *gorm.DB, c *gin.Context) {
	var user models.User
	if err := db.Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to get users"})
	}
	c.JSON(http.StatusOK, gin.H{"users": user})
}

func GetUserById(db *gorm.DB, c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "userid/user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

	}
}
