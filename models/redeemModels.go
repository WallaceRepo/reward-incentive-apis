package model

import "gorm.io/gorm"

// Define the struct for the redeemActivityType
type Redeem struct {
	ID          uint    `gorm:"primaryKey"`
	Type        string  // Type of redemption (e.g., "purchase", "free_product", "free_delivery", etc.)
	Description string  // Description of the redemption option
	MaxValueUSD float64 // Maximum USD value for the redemption option
}

// Function to initialize redeemActivityType with default data
func InitializeRedeem(db *gorm.DB) error {
	// Default data for redeemActivityType
	defaultData := []Redeem{
		{
			Type:        "purchase",
			Description: "Redeem for credits towards a purchase order",
			MaxValueUSD: 100.0,
		},
		{
			Type:        "free_product",
			Description: "Redeem for a free product (up to certain USD value)",
			MaxValueUSD: 50.0,
		},
		{
			Type:        "free_delivery",
			Description: "Redeem for free delivery (up to certain USD value)",
			MaxValueUSD: 20.0,
		},
		{
			Type:        "merchandise",
			Description: "Purchase Gaze-branded merchandise",
			MaxValueUSD: 75.0,
		},
		{
			Type:        "giveaway",
			Description: "Redeem for giveaway entries",
			MaxValueUSD: 10.0,
		},
		{
			Type:        "donate",
			Description: "Donate to a charity",
			MaxValueUSD: 50.0,
		},
		{
			Type:        "transfer",
			Description: "Transfer to another Gaze customer account",
			MaxValueUSD: 200.0,
		},
		{
			Type:        "cash_out",
			Description: "Cash-out points",
			MaxValueUSD: 200.0,
		},
	}

	// Insert default data
	for _, data := range defaultData {
		result := db.FirstOrCreate(&data, Redeem{Type: data.Type})
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
