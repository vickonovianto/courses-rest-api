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

type courseDelivery struct {
	courseUsecase model.CourseUsecase
}

type CourseDelivery interface {
	Mount(group *echo.Group)
}

func NewCourseDelivery(courseUsecase model.CourseUsecase) CourseDelivery {
	return &courseDelivery{courseUsecase: courseUsecase}
}

func (p *courseDelivery) Mount(group *echo.Group) {
	group.POST("", p.StoreCourseHandler)
	group.GET("", p.FetchCourseHandler)
	group.GET("/:id", p.DetailCourseHandler)
	group.PATCH("/:id", p.EditCourseHandler)
	group.DELETE("/:id", p.DeleteCourseHandler)
}

func (p *courseDelivery) StoreCourseHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	var req request.CourseRequest
	if err := c.Bind(&req); err != nil {
		return helper.ResponseValidationErrorJson(c, "Error binding struct", err.Error())
	}
	if err := req.Validate(); err != nil {
		errVal := err.(validation.Errors)
		return helper.ResponseValidationErrorJson(c, "Error validation", errVal)
	}
	course, err := p.courseUsecase.StoreCourse(ctx, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	return helper.ResponseSuccessJson(c, "success", course)
}

func (p *courseDelivery) FetchCourseHandler(c echo.Context) error {
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
	courses, err := p.courseUsecase.FetchCourse(ctx, limitInt, offsetInt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return helper.ResponseSuccessJson(c, "success", courses)
}

func (p *courseDelivery) DetailCourseHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	course, err := p.courseUsecase.GetByID(ctx, idInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	return helper.ResponseSuccessJson(c, "success", course)

}

func (p *courseDelivery) EditCourseHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	var req request.CourseRequest
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
	course, err := p.courseUsecase.EditCourse(ctx, idInt, &req)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}
	return helper.ResponseSuccessJson(c, "success", course)
}

func (p *courseDelivery) DeleteCourseHandler(c echo.Context) error {
	if err := helper.CheckTokenClaim(c); err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnauthorized, err)
	}
	ctx := c.Request().Context()
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusBadRequest, err)
	}
	err = p.courseUsecase.DestroyCourse(ctx, idInt)
	if err != nil {
		return helper.ResponseErrorJson(c, http.StatusUnprocessableEntity, err)
	}
	return helper.ResponseSuccessJson(c, "success", "")
}
