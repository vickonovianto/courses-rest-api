package model

import (
	"context"
	"courses-api/request"
)

type (
	Course struct {
		ID               int64           `json:"id"`
		Title            string          `json:"title"`
		CourseCategoryID int64           `json:"course_category_id"`
		CourseCategory   *CourseCategory `json:"course_category"`
	}

	CourseRepository interface {
		Create(ctx context.Context, course *Course) (*Course, error)
		Fetch(ctx context.Context, limit, offset int) ([]*Course, error)
		FindByID(ctx context.Context, id int64) (*Course, error)
		UpdateByID(ctx context.Context, id int64, course *Course) (*Course, error)
		Delete(ctx context.Context, id int64) error
	}

	CourseUsecase interface {
		StoreCourse(ctx context.Context, req *request.CourseRequest) (*Course, error)
		FetchCourse(ctx context.Context, limit, offset int) ([]*Course, error)
		GetByID(ctx context.Context, id int64) (*Course, error)
		EditCourse(ctx context.Context, id int64, req *request.CourseRequest) (*Course, error)
		DestroyCourse(ctx context.Context, id int64) error
	}
)

// override gorm table name
func (Course) TableName() string {
	return "courses"
}
