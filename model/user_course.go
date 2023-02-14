package model

import (
	"context"
	"courses-api/request"
)

type (
	UserCourse struct {
		ID       int64   `json:"id"`
		UserID   int64   `json:"user_id"`
		User     *User   `json:"user"`
		CourseID int64   `json:"course_id"`
		Course   *Course `json:"course"`
	}

	UserCourseRepository interface {
		Create(ctx context.Context, userCourse *UserCourse) (*UserCourse, error)
		Fetch(ctx context.Context, limit, offset int) ([]*UserCourse, error)
		FindByID(ctx context.Context, id int64) (*UserCourse, error)
		UpdateByID(ctx context.Context, id int64, userCourse *UserCourse) (*UserCourse, error)
		Delete(ctx context.Context, id int64) error
	}

	UserCourseUsecase interface {
		StoreUserCourse(ctx context.Context, req *request.UserCourseRequest) (*UserCourse, error)
		FetchUserCourse(ctx context.Context, limit, offset int) ([]*UserCourse, error)
		GetByID(ctx context.Context, id int64) (*UserCourse, error)
		EditUserCourse(ctx context.Context, id int64, req *request.UserCourseRequest) (*UserCourse, error)
		DestroyUserCourse(ctx context.Context, id int64) error
	}
)

// override gorm table name
func (UserCourse) TableName() string {
	return "user_courses"
}
