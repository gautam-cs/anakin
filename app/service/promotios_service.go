package service

import (
	"gautam/server/app/db/promotions_repo"
	"gautam/server/app/db/users_repo"
	"gautam/server/app/models"
	"gautam/server/app/resource/query"
	"gautam/server/app/utils"
	"github.com/phuslu/log"
	"time"
)

func RunPromotions(user string, request *query.PromotionsRequest) error {
	isPromotionActive := 0

	if request.StartDate == nil {
		isPromotionActive = 1
	}

	promotionObj := &models.Promotions{
		UUID:         utils.UUIDV4(),
		RetailerID:   request.RetailerID,
		ProductID:    request.ProductID,
		IsActive:     isPromotionActive,
		Discount:     request.Discount,
		CreatedDate:  models.DateTime(time.Now()),
		ModifiedDate: models.DateTime(time.Now()),
	}
	if request.StartDate != nil {
		dt, err := time.Parse(string(models.TimeLayoutDateTimeSQL), *request.StartDate)
		if err != nil {
			return nil
		}
		startDate := models.DateTime(dt)
		promotionObj.StartTime = &startDate
	}
	if request.EndDate != nil {
		dt, err := time.Parse(string(models.TimeLayoutDateTimeSQL), *request.EndDate)
		if err != nil {
			return nil
		}
		endDate := models.DateTime(dt)
		promotionObj.EndTime = &endDate
	}

	if err := promotions_repo.Create(promotionObj); err != nil {
		return err
	}

	users, err := users_repo.FindAllUserWithoutUsername(user)
	if err != nil {
		return err
	}
	go sendPromotionEmails(users)

	return nil
}

func sendPromotionEmails(users []*models.Users) error {
	emails := make([]string, 0)
	for _, item := range users {
		emails = append(emails, item.Email)
	}
	if err := sendEmail("test@gmail.com", "password", emails, "", ""); err != nil {
		log.Error().Err(err).Msg("sendEmail failed")
	}
	return nil
}
