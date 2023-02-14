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

type courseCategoryDelivery struct {
	courseCategoryUsecase model.CourseCategoryUsecase
}

type CourseCategoryDelivery interface {
	Mount(group *echo.Group)
}

func NewCourseCategoryDelivery(courseCategoryUsecase model.CourseCategoryUsecase) CourseCategoryDelivery {
	return &courseCategoryDelivery{courseCategoryUsecase: courseCategoryUsecase}
}

func (p *courseCategoryDelivery) Mount(group *echo.Group) {
	group.POST("", p.StoreCourseCategoryHandler)
	group.GET("", p.FetchCourseCategoryHandler)
	group.GET("/:id", p.DetailCourseCategoryHandler)
	group.PATCH("/:id", p.EditCourseCategoryHandler)
	group.DELETE("/:id", p.DeleteCourseCategoryHandler)
}

func (p *courseCategoryDelivery) StoreCourseCategoryHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	var req request.CourseCategoryRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}
	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	courseCategory, err := p.courseCategoryUsecase.StoreCourseCategory(ctx, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	return helper.ResponseSuccessJson(c, "success", courseCategory)
}

func (p *courseCategoryDelivery) FetchCourseCategoryHandler(c echo.Context) error {
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
	courseCategories, err := p.courseCategoryUsecase.FetchCourseCategory(ctx, limitInt, offsetInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return helper.ResponseSuccessJson(c, "success", courseCategories)
}

func (p *courseCategoryDelivery) DetailCourseCategoryHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	courseCategory, err := p.courseCategoryUsecase.GetByID(ctx, idInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	return helper.ResponseSuccessJson(c, "success", courseCategory)
}

func (p *courseCategoryDelivery) EditCourseCategoryHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	var req request.CourseCategoryRequest
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
	courseCategory, err := p.courseCategoryUsecase.EditCourseCategory(ctx, idInt, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}
	return helper.ResponseSuccessJson(c, "success", courseCategory)
}

func (p *courseCategoryDelivery) DeleteCourseCategoryHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	err = p.courseCategoryUsecase.DestroyCourseCategory(ctx, idInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}
	return helper.ResponseSuccessJson(c, "success", "")
}
