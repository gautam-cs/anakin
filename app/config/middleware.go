package config

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/phuslu/log"
)

func EchoJWTMiddleWare() echo.MiddlewareFunc {
	if IsDebugEnv() {
		return dummyJWTMiddleware()
	}

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    []byte(accountJWTKey()),
	})
}

func dummyJWTMiddleware() echo.MiddlewareFunc {
	log.Info().Msg("setting up dummy jwt middleware")

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			val := c.Request().Header.Get("Authorization")
			if val == "" {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "auth header missing",
					Internal: nil,
				}
			}

			parts := strings.SplitN(val, " ", 2)

			if len(parts) != 2 {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "missing bearer token auth",
					Internal: nil,
				}
			}

			token, _, _ := new(jwt.Parser).ParseUnverified(parts[1], jwt.MapClaims{})

			if token != nil {
				c.Set("user", token)
				return next(c)
			}

			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  "missing valid jwt token claims",
				Internal: nil,
			}
		}
	}
}
