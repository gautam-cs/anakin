package users_repo

import (
	"gautam/server/app/config"
	"gautam/server/app/models"
)

func Create(createDate *models.Users) error {
	err := config.WriteDB().
		Model(&models.Users{}).
		Create(createDate).Error
	return err
}

func FindUserByUsername(username string) (*models.Users, error) {
	user := new(models.Users)
	if err := config.ReadDB().
		Where(&models.Users{Username: username}).
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func FindAllUserWithoutUsername(username string) ([]*models.Users, error) {
	users := make([]*models.Users, 0)
	if e := config.ReadDB().
		Model(&models.Users{}).
		Where("username != ?", username).
		Find(&users).Error; e != nil {
		return users, e
	}
	return users, nil
}
