package main

import (
	"github.com/labstack/echo/v4"

	"gautam/server/app/config"
	"gautam/server/app/resource/apis"
	"gautam/server/app/service"
)

func addRoutes(e *echo.Echo) {
	jwtMiddleWare := config.EchoJWTMiddleWare()

	v2 := e.Group("/anakin")
	v2.GET("/status", service.Status)

	addAccountRoutes(v2, jwtMiddleWare)
	addGuestRoutes(v2)
}

func addGuestRoutes(router *echo.Group) {
	router.POST("/signup", apis.SignUp)
	router.POST("/login", apis.Login)
	router.GET("/products/retailers", apis.ProductsByRetailers)
}

func addAccountRoutes(router *echo.Group, jwtMiddleware echo.MiddlewareFunc) {
	router.POST("/run_promotion", apis.RunPromotion, jwtMiddleware)

}
