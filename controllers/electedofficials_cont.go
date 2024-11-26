package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"muyi2905/civic/backend/models"
	"net/http"
)

// CreateOfficial handles creating a new elected official
func CreateOfficial(db *gorm.DB, c *gin.Context) {
	var official models.ElectedOfficial
	if err := c.ShouldBindJSON(&official); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&official).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create official"})
		return
	}

	c.JSON(http.StatusCreated, official)
}

// GetOfficials fetches all elected officials
func GetOfficials(db *gorm.DB, c *gin.Context) {
	var officials []models.ElectedOfficial
	if err := db.Find(&officials).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch officials"})
		return
	}

	c.JSON(http.StatusOK, officials)
}

// GetOfficial fetches a single elected official by ID
func GetOfficial(db *gorm.DB, c *gin.Context) {
	id := c.Param("id")

	var official models.ElectedOfficial
	if err := db.First(&official, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Official not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch official"})
		return
	}

	c.JSON(http.StatusOK, official)
}

// UpdateOfficial updates an elected official's details
func UpdateOfficial(db *gorm.DB, c *gin.Context) {
	id := c.Param("id")
	var updatedOfficial models.ElectedOfficial

	if err := c.ShouldBindJSON(&updatedOfficial); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var official models.ElectedOfficial
	if err := db.First(&official, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Official not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch official"})
		return
	}

	// Update official's fields
	official.Name = updatedOfficial.Name
	official.Office = updatedOfficial.Office
	official.Party = updatedOfficial.Party
	official.State = updatedOfficial.State
	official.District = updatedOfficial.District

	if err := db.Save(&official).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update official"})
		return
	}

	c.JSON(http.StatusOK, official)
}

// DeleteOfficial deletes an elected official by ID
func DeleteOfficial(db *gorm.DB, c *gin.Context) {
	id := c.Param("id")

	var official models.ElectedOfficial
	if err := db.First(&official, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Official not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch official"})
		return
	}

	if err := db.Delete(&official).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete official"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Official deleted successfully"})
}
