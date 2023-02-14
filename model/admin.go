package model

import (
	"context"
	"courses-api/request"

	"github.com/golang-jwt/jwt/v4"
)

type (
	Admin struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AdminRepository interface {
		Create(ctx context.Context, admin *Admin) (*Admin, error)
		Get(ctx context.Context) (*Admin, error)
	}

	AdminUsecase interface {
		RegisterAdmin(ctx context.Context, req request.AdminRegisterRequest) (*Admin, error)
		LoginAdmin(ctx context.Context, req request.AdminLoginRequest) (*jwt.Token, error)
	}

	CustomClaims struct {
		IsAdmin bool `json:"admin"`
		jwt.RegisteredClaims
	}
)

// override gorm table name
func (Admin) TableName() string {
	return "admin"
}
