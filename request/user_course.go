package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	UserCourseRequest struct {
		UserID   int64 `json:"user_id"`
		CourseID int64 `json:"course_id"`
	}
)

func (req UserCourseRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.UserID, validation.Required),
		validation.Field(&req.CourseID, validation.Required),
	)
}
