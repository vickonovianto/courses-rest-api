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

type userCourseDelivery struct {
	userCourseUsecase model.UserCourseUsecase
}

type UserCourseDelivery interface {
	Mount(group *echo.Group)
}

func NewUserCourseDelivery(userCourseUsecase model.UserCourseUsecase) UserCourseDelivery {
	return &userCourseDelivery{userCourseUsecase: userCourseUsecase}
}

func (p *userCourseDelivery) Mount(group *echo.Group) {
	group.POST("", p.StoreUserCourseHandler)
	group.GET("", p.FetchUserCourseHandler)
	group.GET("/:id", p.DetailUserCourseHandler)
	group.PATCH("/:id", p.EditUserCourseHandler)
	group.DELETE("/:id", p.DeleteUserCourseHandler)
}

func (p *userCourseDelivery) StoreUserCourseHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	var req request.UserCourseRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}
	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	userCourse, err := p.userCourseUsecase.StoreUserCourse(ctx, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	return helper.ResponseSuccessJson(c, "success", userCourse)
}

func (p *userCourseDelivery) FetchUserCourseHandler(c echo.Context) error {
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
	userCourses, err := p.userCourseUsecase.FetchUserCourse(ctx, limitInt, offsetInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return helper.ResponseSuccessJson(c, "success", userCourses)
}

func (p *userCourseDelivery) DetailUserCourseHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	userCourse, err := p.userCourseUsecase.GetByID(ctx, idInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	return helper.ResponseSuccessJson(c, "success", userCourse)

}

func (p *userCourseDelivery) EditUserCourseHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	var req request.UserCourseRequest
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
	userCourse, err := p.userCourseUsecase.EditUserCourse(ctx, idInt, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}
	return helper.ResponseSuccessJson(c, "success", userCourse)
}

func (p *userCourseDelivery) DeleteUserCourseHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	err = p.userCourseUsecase.DestroyUserCourse(ctx, idInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}
	return helper.ResponseSuccessJson(c, "success", "")
}
