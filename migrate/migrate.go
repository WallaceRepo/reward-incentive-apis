package migrate

import (
	"github.com/loyaltiReward-api/initializers"
	model "github.com/loyaltiReward-api/models"
)

func Migrate() {
	initializers.DB.AutoMigrate(&model.Redeem{})
	initializers.DB.AutoMigrate(&model.Reward{})
	initializers.DB.AutoMigrate(&model.ShopperActivity{})
	model.AddInitialData(initializers.DB)
	model.InitializeRedeem(initializers.DB)
}
