package handler

import (
	"net/http"
	"strconv"

	model "github.com/loyaltiReward-api/models"
	"github.com/gin-gonic/gin"
	"github.com/loyaltiReward-api/initializers"
)

// CreateRedeem creates a new redeem activity type.
func CreateRedeem(c *gin.Context) {
	var redeem model.Redeem
	if err := c.ShouldBindJSON(&redeem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	err := initializers.DB.Create(&redeem).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create redeem"})
		return
	}

	c.JSON(http.StatusCreated, redeem)
}

// UpdateRedeem updates an existing redeem activity type by ID.
func UpdateRedeem(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var redeem model.Redeem
	if err := initializers.DB.First(&redeem, uint(uintID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Redeem not found"})
		return
	}

	if err := c.ShouldBindJSON(&redeem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	err = initializers.DB.Save(&redeem).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update redeem"})
		return
	}

	c.JSON(http.StatusOK, redeem)
}

// DeleteRedeem deletes a redeem by ID.
func DeleteRedeem(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = initializers.DB.Delete(&model.Redeem{}, uint(uintID)).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete redeem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Redeem deleted successfully"})
}

// GetAllRedeems function to get all redeem activity types
func GetAllRedeems(c *gin.Context) {
	var redeems []model.Redeem

	if err := initializers.DB.Find(&redeems).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch redeems"})
		return
	}

	c.JSON(http.StatusOK, redeems)
}

// GetRedeem function to get a redeem by ID
func GetRedeem(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var redeem model.Redeem
	if err := initializers.DB.First(&redeem, uint(uintID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Redeem not found"})
		return
	}

	c.JSON(http.StatusOK, redeem)
}
