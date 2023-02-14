package repository

import (
	"context"
	"courses-api/config"
	"courses-api/model"
)

type userRepository struct {
	Cfg config.Config
}

func NewUserRepository(cfg config.Config) model.UserRepository {
	return &userRepository{Cfg: cfg}
}

func (u *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := u.Cfg.Database().WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.User, error) {
	var data []*model.User
	if err := u.Cfg.Database().WithContext(ctx).
		Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userRepository) FindByID(ctx context.Context, id int64) (*model.User, error) {
	user := new(model.User)
	if err := u.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) UpdateByID(ctx context.Context, id int64, user *model.User) (*model.User, error) {
	if err := u.Cfg.Database().WithContext(ctx).
		Model(&model.User{ID: id}).Updates(user).Find(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepository) Delete(ctx context.Context, id int64) error {
	_, err := u.FindByID(ctx, id)
	if err != nil {
		return err
	}
	res := u.Cfg.Database().WithContext(ctx).
		Delete(&model.User{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
