package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type (
	AdminRegisterRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AdminLoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

func (req AdminRegisterRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Name, validation.Required, validation.Length(1, 255)),
		validation.Field(&req.Email, validation.Required, is.Email, validation.Length(3, 255)),
		validation.Field(&req.Password, validation.Required, validation.Length(6, 255)),
	)
}

func (req AdminLoginRequest) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Email, validation.Required, is.Email, validation.Length(3, 255)),
		validation.Field(&req.Password, validation.Required, validation.Length(6, 255)),
	)
}
