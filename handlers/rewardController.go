package handler

import (
	"net/http"
	"strconv"

	model "github.com/loyaltiReward-api/models"
	"github.com/gin-gonic/gin"
	"github.com/loyaltiReward-api/initializers"
)

// CreateReward creates a new reward activity type.
func CreateReward(c *gin.Context) {
	var reward model.Reward
	if err := c.ShouldBindJSON(&reward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	err := initializers.DB.Create(&reward).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reward"})
		return
	}

	c.JSON(http.StatusCreated, reward)
}

// UpdateReward updates an existing reward activity type by ID.
func UpdateReward(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var reward model.Reward
	if err := initializers.DB.First(&reward, uint(uintID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reward not found"})
		return
	}

	if err := c.ShouldBindJSON(&reward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data format"})
		return
	}

	err = initializers.DB.Save(&reward).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reward"})
		return
	}

	c.JSON(http.StatusOK, reward)
}

// DeleteReward deletes a reward by ID.
func DeleteReward(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = initializers.DB.Delete(&model.Reward{}, uint(uintID)).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete reward"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reward deleted successfully"})
}

// GetAllRewards function to get all reward activity types
func GetAllRewards(c *gin.Context) {
	var rewards []model.Reward

	if err := initializers.DB.Find(&rewards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch rewards"})
		return
	}

	c.JSON(http.StatusOK, rewards)
}

// GetReward function to get a reward by ID
func GetReward(c *gin.Context) {
	id := c.Param("id")
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var reward model.Reward
	if err := initializers.DB.First(&reward, uint(uintID)).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reward not found"})
		return
	}

	c.JSON(http.StatusOK, reward)
}
