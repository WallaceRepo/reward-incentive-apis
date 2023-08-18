package model

import "time"

// Define ShopperActivity struct for the activity information
type ShopperActivity struct {
    ID           uint   `gorm:"primary_key"`
    ShopperGID    uint   // Foreign key to link to the Shopper struct
    Activity string // Indicates the type of activity: "redeem" or "reward"
    PointsBalance      int    // The points associated with the activity, positive for reward and negative for redeem
    TimeStamp    time.Time
}




