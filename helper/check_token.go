package helper

import (
	"courses-api/model"
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func CheckTokenClaim(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*model.CustomClaims)
	isAdmin := claims.IsAdmin
	if isAdmin {
		return nil
	} else {
		return errors.New("invalid token claims")
	}
}
