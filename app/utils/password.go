package utils

import (
	"gautam/server/app/accounts"
	"gautam/server/app/config"
	"gautam/server/app/models"
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
	"net/http"
)

func GeneratePassword(c echo.Context) error {
	userinfo, err := accounts.ValidateAuth(c)
	if err != nil {
		return ErrorResponse(c, http.StatusUnauthorized, err)
	}

	if userinfo.Role != accounts.IUserRoleTypeSUser {
		return ErrorResponse(c, http.StatusForbidden, err)
	}

	password := UUIDV4()

	encoded, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		log.Error().Err(err).Msg("GeneratePassword")
		return ErrorResponse(c, http.StatusInternalServerError, err)
	}

	item := &models.Users{
		Username: "ActiveCampaignUser",
		Password: string(encoded),
	}

	err = config.WriteDB().
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "username"}},
			DoUpdates: clause.AssignmentColumns([]string{"hash_key"}),
		}).
		Create(item).Error

	if err != nil {
		log.Error().Err(err).Msg("GeneratePassword")
		return ErrorResponse(c, http.StatusInternalServerError, err)
	}

	result := map[string]interface{}{
		"access_key": password,
	}

	return c.JSON(http.StatusOK, result)
}

func GenerateAuthToken(c echo.Context) error {
	username, password, ok := c.Request().BasicAuth()

	if !ok {
		return ErrorResponsef(c, http.StatusBadRequest, "Authorization missing")
	}

	if username != "ActiveCampaignUser" {
		log.Error().Msgf("Active campaign login failed for username: %s, password: %s", username, password)
		return ErrorResponsef(c, http.StatusBadRequest, "User/Password do not match")
	}

	dbItem := new(models.Users)

	err := config.ReadDB().
		Where(&models.Users{Username: username}).
		First(dbItem).Error

	if err != nil {
		log.Error().Msgf("Active campaign login failed for username: %s, password: %s", username, password)
		return ErrorResponsef(c, http.StatusBadRequest, "User/Password do not match")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbItem.Password), []byte(password)); err != nil {
		log.Error().Msgf("Active campaign login failed for username: %s, password: %s", username, password)
		return ErrorResponsef(c, http.StatusBadRequest, "User/Password do not match")
	}

	expireAfterSeconds := int64(86400)
	token, err := accounts.MakeJWTTokenWithExpiry(username, accounts.IUserRoleTypeRoot, expireAfterSeconds)
	if err != nil {
		log.Error().Err(err).Msg("error generating token")
		return ErrorResponsef(c, http.StatusInternalServerError, "error generating token")
	}

	result := map[string]interface{}{
		"token_type":   "Bearer",
		"expires_in":   expireAfterSeconds,
		"access_token": token,
	}

	return c.JSON(http.StatusOK, result)
}
