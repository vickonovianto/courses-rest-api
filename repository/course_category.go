package repository

import (
	"context"
	"courses-api/config"
	"courses-api/model"
)

type courseCategoryRepository struct {
	Cfg config.Config
}

func NewCourseCategoryRepository(cfg config.Config) model.CourseCategoryRepository {
	return &courseCategoryRepository{Cfg: cfg}
}

func (c *courseCategoryRepository) Create(ctx context.Context, courseCategory *model.CourseCategory) (*model.CourseCategory, error) {
	if err := c.Cfg.Database().WithContext(ctx).Create(&courseCategory).Error; err != nil {
		return nil, err
	}
	return courseCategory, nil
}

func (c *courseCategoryRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.CourseCategory, error) {
	var data []*model.CourseCategory

	if err := c.Cfg.Database().WithContext(ctx).
		Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (c *courseCategoryRepository) FindByID(ctx context.Context, id int64) (*model.CourseCategory, error) {
	courseCategory := new(model.CourseCategory)

	if err := c.Cfg.Database().
		WithContext(ctx).
		Where("id = ?", id).
		First(courseCategory).Error; err != nil {
		return nil, err
	}
	return courseCategory, nil
}

func (c *courseCategoryRepository) UpdateByID(ctx context.Context, id int64, courseCategory *model.CourseCategory) (*model.CourseCategory, error) {
	if err := c.Cfg.Database().WithContext(ctx).
		Model(&model.CourseCategory{ID: id}).Updates(courseCategory).Find(courseCategory).Error; err != nil {
		return nil, err
	}
	return courseCategory, nil
}

func (c *courseCategoryRepository) Delete(ctx context.Context, id int64) error {
	_, err := c.FindByID(ctx, id)
	if err != nil {
		return err
	}

	res := c.Cfg.Database().WithContext(ctx).
		Delete(&model.CourseCategory{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
