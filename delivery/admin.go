package delivery

import (
	"courses-api/helper"
	"courses-api/model"
	"courses-api/request"
	"net/http"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

type adminDelivery struct {
	adminUsecase model.AdminUsecase
}

type AdminDelivery interface {
	Mount(group *echo.Group)
}

func NewAdminDelivery(adminUsecase model.AdminUsecase) AdminDelivery {
	return &adminDelivery{adminUsecase: adminUsecase}
}

func (a *adminDelivery) Mount(group *echo.Group) {
	group.POST("/register", a.RegisterAdminHandler)
	group.POST("/login", a.LoginAdminHandler)
}

func (a *adminDelivery) RegisterAdminHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.AdminRegisterRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}
	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	admin, err := a.adminUsecase.RegisterAdmin(ctx, req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	return helper.ResponseSuccessJson(c, "success", admin)
}

func (a *adminDelivery) LoginAdminHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req request.AdminLoginRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}
	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	token, err := a.adminUsecase.LoginAdmin(ctx, req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	// Generate encoded token and send it as response.
	signedToken, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return err
	}
	return helper.ResponseSuccessJson(c, "success", echo.Map{
		"token": signedToken,
	})
}
