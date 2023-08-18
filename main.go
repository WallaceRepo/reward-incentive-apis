package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/loyaltiReward-api/initializers"
	"github.com/loyaltiReward-api/migrate"
	model "github.com/loyaltiReward-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	migrate.Migrate()

}

func main() {
	fmt.Println("Hello")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1/incentive")
	{
		// // Reward default data endpoints
		// v1.POST("/rewards_lookup", handler.CreateReward)
		// v1.PUT("/rewards_lookup/:id", handler.UpdateReward)
		// v1.DELETE("/rewards_lookup/:id", handler.DeleteReward)
		// v1.GET("/rewards_lookup", handler.GetAllRewards)
		// v1.GET("/rewards_lookup/:id", handler.GetReward)

		// // Redeem default data endpoints
		// v1.POST("/redeems_lookup", handler.CreateRedeem)
		// v1.PUT("/redeems_lookup/:id", handler.UpdateRedeem)
		// v1.DELETE("/redeems_lookup/:id", handler.DeleteRedeem)
		// v1.GET("/redeems_lookup", handler.GetAllRedeems)
		// v1.GET("/redeems_lookup/:id", handler.GetRedeem)

		// // Shopper activities
            // for testing via shopper-seller-api
		v1.GET("/accounts/shopper", SaveShopper) // creates Shopper account on Shopper-Seller-API

		//reward APIs
		v1.POST("/reward/create-new-shopper", CreateShopperIncentive) // create new shopper incentive
		v1.PUT("/reward/:id")                                         // add/update shopper incentive
		v1.GET("/reward/:id")                                         // get shopper's points

		// redeem
		v1.PUT("/redeem/:id") // deletes/uupdates point of shopper on Shopper Intensive

	}

	router.Run()
}

func CreateShopperIncentive(c *gin.Context) {
	var payload struct {
		ShopperID    uint   `json:"shopperID"`
		ActivityType string `json:"activitytype"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the reward points from the Reward table based on the activity type
	var reward model.Reward
	if err := initializers.DB.Where("Name = ?", payload.ActivityType).First(&reward).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Reward not found"})
		return
	}

	// Create a new shopper activity
	shopperActivity := model.ShopperActivity{
		ShopperGID:    payload.ShopperID,
		Activity:      payload.ActivityType,
		PointsBalance: reward.Points, // Initialize with reward points
		TimeStamp:     time.Now(),
	}
	// Insert the shopper activity into the database
	if err := initializers.DB.Create(&shopperActivity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create shopper activity"})
		return
	}

	c.JSON(http.StatusOK, shopperActivity)
}

// Test func; to send noitification from shopper-seller-api to reward-loyalty-api: v1.GET("/accounts/shopper", SaveShopper)
func SaveShopper(c *gin.Context) {
	// Hardcoded shopper payload
	shopperPayload := []byte(`{"ShopperID": 1, "ActivityType": "create-new-shopper"}`)

	// Notify the loyalty-reward-api about the new shopper
	receiverURL := "http://localhost:3000/api/v1/incentive/reward/create-new-shopper"

	// Send HTTP request to the receiver microservice
	_, err := http.Post(receiverURL, "application/json", bytes.NewBuffer(shopperPayload))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to notify receiver service"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shopper saved successfully"})
}
