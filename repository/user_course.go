package repository

import (
	"context"
	"courses-api/config"
	"courses-api/model"
)

type userCourseRepository struct {
	Cfg config.Config
}

func NewUserCourseRepository(cfg config.Config) model.UserCourseRepository {
	return &userCourseRepository{Cfg: cfg}
}

func (u *userCourseRepository) Create(ctx context.Context, userCourse *model.UserCourse) (*model.UserCourse, error) {
	if err := u.Cfg.Database().WithContext(ctx).Create(&userCourse).Error; err != nil {
		return nil, err
	}
	if err := u.Cfg.Database().
		WithContext(ctx).
		Preload("User").
		Preload("Course").
		Where("id = ?", userCourse.ID).
		First(userCourse).Error; err != nil {
		return nil, err
	}
	return userCourse, nil
}

func (u *userCourseRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.UserCourse, error) {
	var data []*model.UserCourse
	if err := u.Cfg.Database().WithContext(ctx).Preload("User").Preload("Course").
		Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userCourseRepository) FindByID(ctx context.Context, id int64) (*model.UserCourse, error) {
	userCourse := new(model.UserCourse)
	if err := u.Cfg.Database().
		WithContext(ctx).
		Preload("User").
		Preload("Course").
		Where("id = ?", id).
		First(userCourse).Error; err != nil {
		return nil, err
	}
	return userCourse, nil
}

func (u *userCourseRepository) UpdateByID(ctx context.Context, id int64, userCourse *model.UserCourse) (*model.UserCourse, error) {
	if err := u.Cfg.Database().WithContext(ctx).
		Model(&model.UserCourse{ID: id}).Updates(userCourse).Find(userCourse).Error; err != nil {
		return nil, err
	}
	if err := u.Cfg.Database().
		WithContext(ctx).
		Preload("User").
		Preload("Course").
		Where("id = ?", userCourse.ID).
		First(userCourse).Error; err != nil {
		return nil, err
	}
	return userCourse, nil
}

func (u *userCourseRepository) Delete(ctx context.Context, id int64) error {
	_, err := u.FindByID(ctx, id)
	if err != nil {
		return err
	}
	res := u.Cfg.Database().WithContext(ctx).
		Delete(&model.UserCourse{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
