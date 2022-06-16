package promotions_repo

import (
	"gautam/server/app/config"
	"gautam/server/app/models"
)

func Create(createDate *models.Promotions) error {
	err := config.WriteDB().
		Model(&models.Promotions{}).
		Create(createDate).Error
	return err
}
