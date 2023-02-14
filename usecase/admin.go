package usecase

import (
	"context"
	"courses-api/model"
	"courses-api/request"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type adminUsecase struct {
	adminRepository model.AdminRepository
}

func NewAdminUsecase(adminRepository model.AdminRepository) model.AdminUsecase {
	return &adminUsecase{adminRepository: adminRepository}
}

func (a *adminUsecase) RegisterAdmin(ctx context.Context, req request.AdminRegisterRequest) (*model.Admin, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(bytes)

	admin, err := a.adminRepository.Create(ctx, &model.Admin{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})

	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (a *adminUsecase) LoginAdmin(ctx context.Context, req request.AdminLoginRequest) (*jwt.Token, error) {
	admin, err := a.adminRepository.Get(ctx)
	if err != nil {
		return nil, err
	}
	if admin.Email == req.Email {
		hashedPassword := admin.Password
		err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
		if err != nil {
			return nil, errors.New("incorrect email or password")
		}

		// Set custom claims
		claims := &model.CustomClaims{
			IsAdmin: true,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		// Create token with claims
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		return token, nil
	}
	return nil, errors.New("incorrect email or password")
}
