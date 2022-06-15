package service

import (
	"gautam/server/app/accounts"
	"gautam/server/app/config"
	"gautam/server/app/models"
	"gautam/server/app/resource/query"

	"github.com/phuslu/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

func SignUp(requestData *query.SignUpRequest) error {

	encoded, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), 4)
	if err != nil {
		log.Error().Err(err).Msg("GeneratePassword")
		return err
	}

	item := &models.Users{
		Username:  requestData.UserName,
		Password:  string(encoded),
		Email:     requestData.Email,
		FirstName: requestData.FirstName,
	}
	if requestData.LastName != nil {
		item.LastName = *requestData.LastName
	}

	err = config.WriteDB().
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "username"}},
			DoUpdates: clause.AssignmentColumns([]string{"password"}),
		}).
		Create(item).Error

	if err != nil {
		log.Error().Err(err).Msg("GeneratePassword")
		return err
	}

	return nil

}

func Login(requestData *query.LoginRequest) (*map[string]interface{}, error) {

	dbItem := new(models.Users)

	err := config.ReadDB().
		Where(&models.Users{Username: requestData.UserName}).
		First(dbItem).Error

	if err != nil {
		log.Error().Msgf("Active campaign login failed for username: %s, password: %s", requestData.UserName, requestData.Password)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbItem.Password), []byte(requestData.Password)); err != nil {
		log.Error().Msgf("Active campaign login failed for username: %s, password: %s", requestData.UserName, requestData.Password)
		return nil, err
	}

	expireAfterSeconds := int64(86400)
	token, err := accounts.MakeJWTTokenWithExpiry(requestData.UserName, accounts.IUserRoleTypeRoot, expireAfterSeconds)
	if err != nil {
		log.Error().Err(err).Msg("error generating token")
		return nil, err
	}

	result := map[string]interface{}{
		"token_type":   "Bearer",
		"expires_in":   expireAfterSeconds,
		"access_token": token,
	}

	return &result, nil

}
