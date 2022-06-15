package service

import (
	"gautam/server/app/config"
	"gautam/server/app/models"
	"gautam/server/app/resource/query"
	"gautam/server/app/utils"
	"github.com/phuslu/log"
	"time"
)

func RunPromotions(request *query.PromotionsRequest) error {
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

	if e := config.WriteDB().Create(promotionObj).Error; e != nil {
		log.Error().Err(e).Msgf("RunPromotions failed for obj : %v", promotionObj)
		return e
	}

	return nil
}
