package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorResponse(c echo.Context, status int, err error) error {
	result := map[string]interface{}{
		"status": "failure",
		"reason": err.Error(),
	}

	return JsonResponse(c, status, result)
}

func ErrorResponsef(c echo.Context, status int, err string) error {
	result := map[string]interface{}{
		"status": "failure",
		"reason": err,
	}

	return JsonResponse(c, status, result)
}

func ErrorResponseI(c echo.Context, status int, err interface{}) error {
	result := map[string]interface{}{
		"status": "failure",
		"reason": err,
	}

	return JsonResponse(c, status, result)
}

func SuccessResponse(c echo.Context) error {
	result := map[string]interface{}{
		"status": "success",
	}

	return JsonResponse(c, http.StatusOK, result)
}

func SuccessResponseWithData(c echo.Context, data interface{}) error {
	result := map[string]interface{}{
		"status": "success",
		"data":   data,
	}

	return JsonResponse(c, http.StatusOK, result)
}

func JsonResponse(c echo.Context, status int, result interface{}) error {
	return c.JSON(status, result)
}
