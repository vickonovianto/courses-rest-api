package delivery

import (
	"courses-api/helper"
	"courses-api/model"
	"courses-api/request"
	"errors"
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
)

type userDelivery struct {
	userUsecase model.UserUsecase
}

type UserDelivery interface {
	Mount(group *echo.Group)
}

func NewUserDelivery(userUsecase model.UserUsecase) UserDelivery {
	return &userDelivery{userUsecase: userUsecase}
}

func (p *userDelivery) Mount(group *echo.Group) {
	group.POST("", p.StoreUserHandler)
	group.GET("", p.FetchUserHandler)
	group.GET("/:id", p.DetailUserHandler)
	group.PATCH("/:id", p.EditUserHandler)
	group.DELETE("/:id", p.DeleteUserHandler)
}

func (p *userDelivery) StoreUserHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	var req request.UserRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}
	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	user, err := p.userUsecase.StoreUser(ctx, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	return helper.ResponseSuccessJson(c, "success", user)
}

func (p *userDelivery) FetchUserHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitInt, offsetInt := -1, -1
	var err error
	if limit != "" {
		limitInt, err = strconv.Atoi(limit)
		if err != nil {
			return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		}
		if limitInt < 1 {
			return helper.ResponseErrorJson(c, http.StatusBadRequest, errors.New("limit must be greater than 0"))
		}
	}
	if offset != "" {
		offsetInt, err = strconv.Atoi(offset)
		if err != nil {
			return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
		}
		if offsetInt < 0 {
			return helper.ResponseErrorJson(c, http.StatusBadRequest, errors.New("offset must not be negative"))
		}
	}
	users, err := p.userUsecase.FetchUser(ctx, limitInt, offsetInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return helper.ResponseSuccessJson(c, "success", users)
}

func (p *userDelivery) DetailUserHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	user, err := p.userUsecase.GetByID(ctx, idInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	return helper.ResponseSuccessJson(c, "success", user)
}

func (p *userDelivery) EditUserHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	var req request.UserRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}
	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	user, err := p.userUsecase.EditUser(ctx, idInt, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}
	return helper.ResponseSuccessJson(c, "success", user)
}

func (p *userDelivery) DeleteUserHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	err = p.userUsecase.DestroyUser(ctx, idInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}
	return helper.ResponseSuccessJson(c, "success", "")
}
