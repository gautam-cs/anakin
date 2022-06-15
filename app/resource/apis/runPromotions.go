package apis

import (
	"gautam/server/app/accounts"
	"gautam/server/app/resource/query"
	"gautam/server/app/service"
	"gautam/server/app/utils"
	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
	"net/http"
)

func RunPromotion(c echo.Context) error {
	userInfo, e := accounts.ValidateAuth(c)
	if e != nil {
		c.NoContent(http.StatusUnauthorized)
	}
	request := new(query.PromotionsRequest)

	if userInfo.Role != accounts.IUserRoleTypeRoot {
		log.Error().Err(e).Msg("only authorised for root user")
		return e
	}

	if e := c.Bind(request); e != nil {
		log.Error().Err(e).Msgf("request binding failed for request %v", request)
		return utils.ErrorResponse(c, http.StatusBadRequest, e)
	}

	if e := service.RunPromotions(request); e != nil {
		log.Error().Err(e).Msgf("service.AddProduct:: Unable to add product to our system")
		return utils.ErrorResponse(c, http.StatusBadGateway, e)
	}
	return utils.SuccessResponse(c)
}
