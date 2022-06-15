package apis

import (
	"gautam/server/app/resource/query"
	"gautam/server/app/service"
	"gautam/server/app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/phuslu/log"
)

func Login(c echo.Context) error {

	request := new(query.LoginRequest)

	if e := c.Bind(request); e != nil {
		return utils.ErrorResponse(c, http.StatusBadRequest, e)
	}

	response, e := service.Login(request)
	if e != nil {
		log.Error().Err(e).Msgf("service.AddProduct:: Unable to add product to our system")
		return utils.ErrorResponse(c, http.StatusBadGateway, e)
	}

	return utils.SuccessResponseWithData(c, response)
}
