package repository

import (
	"context"
	"courses-api/config"
	"courses-api/model"
)

type courseRepository struct {
	Cfg config.Config
}

func NewCourseRepository(cfg config.Config) model.CourseRepository {
	return &courseRepository{Cfg: cfg}
}

func (c *courseRepository) Create(ctx context.Context, course *model.Course) (*model.Course, error) {
	if err := c.Cfg.Database().WithContext(ctx).Create(&course).Error; err != nil {
		return nil, err
	}
	if err := c.Cfg.Database().
		WithContext(ctx).
		Preload("CourseCategory").
		Where("id = ?", course.ID).
		First(course).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func (c *courseRepository) Fetch(ctx context.Context, limit, offset int) ([]*model.Course, error) {
	var data []*model.Course
	if err := c.Cfg.Database().WithContext(ctx).Preload("CourseCategory").
		Limit(limit).Offset(offset).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (c *courseRepository) FindByID(ctx context.Context, id int64) (*model.Course, error) {
	course := new(model.Course)
	if err := c.Cfg.Database().
		WithContext(ctx).
		Preload("CourseCategory").
		Where("id = ?", id).
		First(course).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func (c *courseRepository) UpdateByID(ctx context.Context, id int64, course *model.Course) (*model.Course, error) {
	if err := c.Cfg.Database().WithContext(ctx).
		Model(&model.Course{ID: id}).Updates(course).Find(course).Error; err != nil {
		return nil, err
	}
	if err := c.Cfg.Database().
		WithContext(ctx).
		Preload("CourseCategory").
		Where("id = ?", course.ID).
		First(course).Error; err != nil {
		return nil, err
	}
	return course, nil
}

func (c *courseRepository) Delete(ctx context.Context, id int64) error {
	_, err := c.FindByID(ctx, id)
	if err != nil {
		return err
	}
	res := c.Cfg.Database().WithContext(ctx).
		Delete(&model.Course{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
