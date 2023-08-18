package model

import (
	"log"

	"gorm.io/gorm"
)

// RewardType model to represent each reward activity type and its corresponding points to earn
type Reward struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string // Name of the reward activity type, e.g., "video_review", "purchase", "birthday", etc.
	Points int    // Points earned for the reward activity type
}

// Function to initialize the reward activity types and points to earn in the database
func AddInitialData(db *gorm.DB) {
	// Create reward activity types and set the points to earn for each
	rewards := []Reward {
		{Name: "create-new-shopper", Points: 500 },
		{Name: "video_review", Points: 10},
		{Name: "comments_likes_video_review", Points: 5},
		{Name: "purchase_made_by_others", Points: 20},
		{Name: "liking_a_video", Points: 2},
		{Name: "commenting_on_a_video", Points: 3},
		{Name: "viewing_a_video", Points: 1},
		{Name: "resharing_to_social_media", Points: 8},
		{Name: "product_purchase", Points: 50},
		{Name: "initial_purchase_order", Points: 100},
		{Name: "every_n_purchase_order", Points: 20},
		{Name: "hitting_purchase_order_threshold", Points: 200},
		{Name: "referral_for_new_customer", Points: 150},
		{Name: "birthday", Points: 30},
		{Name: "account_anniversary", Points: 40},
		{Name: "product_searches", Points: 5},
		{Name: "adding_to_shopping_cart", Points: 3},
		{Name: "adding_to_wishlist", Points: 2},
		{Name: "adding_to_collection", Points: 2},
		{Name: "platform_sign_up", Points: 10},
		{Name: "create_account_profile", Points: 5},
		{Name: "daily_login", Points: 7},
		{Name: "active_time_threshold", Points: 15},
		{Name: "playing_mini_games", Points: 20},
		{Name: "new_shopper", Points: 80},
	}

	for _, Reward := range rewards {
		if err := db.Create(&Reward).Error; err != nil {
			log.Fatalf("Failed to initialize reward: %v", err)
		}
	}
}
