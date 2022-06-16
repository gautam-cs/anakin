package service

import (
	"gautam/server/app/accounts"
	"gautam/server/app/db/users_repo"
	"gautam/server/app/models"
	"gautam/server/app/resource/query"
	"gautam/server/app/utils"
	"github.com/phuslu/log"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(requestData *query.SignUpRequest) error {

	encoded, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), 4)
	if err != nil {
		log.Error().Err(err).Msg("GeneratePassword")
		return err
	}

	item := &models.Users{
		UUID:      utils.UUIDV4(),
		Username:  requestData.UserName,
		Password:  string(encoded),
		Email:     requestData.Email,
		FirstName: requestData.FirstName,
	}
	if requestData.LastName != nil {
		item.LastName = *requestData.LastName
	}

	if err := users_repo.Create(item); err != nil {
		return err
	}
	return nil

}

func Login(requestData *query.LoginRequest) (*map[string]interface{}, error) {

	user, err := users_repo.FindUserByUsername(requestData.UserName)

	if err != nil {
		log.Error().Msgf("Active campaign login failed for username: %s, password: %s", requestData.UserName, requestData.Password)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestData.Password)); err != nil {
		log.Error().Msgf("Active campaign login failed for username: %s, password: %s", requestData.UserName, requestData.Password)
		return nil, err
	}

	expireAfterSeconds := int64(36000)
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
