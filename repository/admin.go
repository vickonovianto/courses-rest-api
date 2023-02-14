package repository

import (
	"context"
	"courses-api/config"
	"courses-api/model"
	"errors"

	"gorm.io/gorm"
)

type adminRepository struct {
	Cfg config.Config
}

func NewAdminRepository(cfg config.Config) model.AdminRepository {
	return &adminRepository{Cfg: cfg}
}

func (a *adminRepository) Get(ctx context.Context) (*model.Admin, error) {
	admin := new(model.Admin)

	if err := a.Cfg.Database().WithContext(ctx).First(admin).Error; err != nil {
		return nil, err
	}

	return admin, nil
}

func (a *adminRepository) Create(ctx context.Context, admin *model.Admin) (*model.Admin, error) {

	adminModel := new(model.Admin)

	if err := a.Cfg.Database().WithContext(ctx).Debug().
		First(&adminModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := a.Cfg.Database().WithContext(ctx).Create(&admin).Find(adminModel).Error; err != nil {
				return nil, err
			}

			return adminModel, nil
		}
		return nil, err
	}

	return nil, errors.New("admin already exists")
}
