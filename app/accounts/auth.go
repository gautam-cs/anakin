package accounts

import (
	"fmt"
	"gautam/server/app/config"
	"github.com/labstack/echo/v4"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

func ValidateAuth(c echo.Context) (*UserInfo, error) {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	if userUuid, ok := claims["user_identity"]; ok {

		if roleVal, ok := claims["role"]; ok {
			return onSAuthSuccess(userUuid, roleVal)
		}

		return nil, fmt.Errorf("auth header missing user role")
	}
	return nil, fmt.Errorf("auth header missing valid bearer token")
}

func onSAuthSuccess(userUuid interface{}, role interface{}) (*UserInfo, error) {
	srole := strings.ToLower(role.(string))
	irole := SUserRoleType(srole).IUserRoleType()
	return onAuthSuccess(userUuid, irole)
}

func onAuthSuccess(userUuid interface{}, role IUserRoleType) (*UserInfo, error) {
	info := &UserInfo{
		User: userUuid.(string),
		Role: role,
	}

	return info, nil
}

func MakeJWTToken(user string, role IUserRoleType) (string, error) {
	iat := time.Now().Unix()

	claims := map[string]interface{}{
		"user_identity": user,
		"role":          role.SUserRoleType().String(),
		"iat":           iat,
		"exp":           iat + 3600,
	}

	switch role {
	case IUserRoleTypeRoot:
		claims["root"] = user
	case IUserRoleTypeSUser:
		claims["user"] = user
	default:
		break
	}

	return config.EncodeJWTToken(claims)
}

func MakeJWTTokenWithExpiry(user string, role IUserRoleType, expireAfterSeconds int64) (string, error) {
	iat := time.Now().Unix()

	claims := map[string]interface{}{
		"user_identity": user,
		"role":          role.SUserRoleType().String(),
		"iat":           iat,
		"exp":           iat + expireAfterSeconds,
	}

	switch role {
	case IUserRoleTypeRoot:
		claims["uuid"] = user
	case IUserRoleTypeSUser:
		claims["uuid"] = user
	default:
		break
	}

	return config.EncodeJWTToken(claims)
}
