package model

import (
	"context"
	"courses-api/request"
)

type (
	CourseCategory struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	}

	CourseCategoryRepository interface {
		Create(ctx context.Context, courseCategory *CourseCategory) (*CourseCategory, error)
		Fetch(ctx context.Context, limit, offset int) ([]*CourseCategory, error)
		FindByID(ctx context.Context, id int64) (*CourseCategory, error)
		UpdateByID(ctx context.Context, id int64, courseCategory *CourseCategory) (*CourseCategory, error)
		Delete(ctx context.Context, id int64) error
	}

	CourseCategoryUsecase interface {
		StoreCourseCategory(ctx context.Context, req *request.CourseCategoryRequest) (*CourseCategory, error)
		FetchCourseCategory(ctx context.Context, limit, offset int) ([]*CourseCategory, error)
		GetByID(ctx context.Context, id int64) (*CourseCategory, error)
		EditCourseCategory(ctx context.Context, id int64, req *request.CourseCategoryRequest) (*CourseCategory, error)
		DestroyCourseCategory(ctx context.Context, id int64) error
	}
)

// override gorm table name
func (CourseCategory) TableName() string {
	return "course_categories"
}
