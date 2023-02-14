package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	CourseCategoryRequest struct {
		Name string `json:"name"`
	}
)

func (req CourseCategoryRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required, validation.Length(1, 255)),
	)
}
