package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	CourseRequest struct {
		Title            string `json:"title"`
		CourseCategoryID int64  `json:"course_category_id"`
	}
)

func (req CourseRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Title, validation.Required, validation.Length(1, 255)),
		validation.Field(&req.CourseCategoryID, validation.Required),
	)
}
