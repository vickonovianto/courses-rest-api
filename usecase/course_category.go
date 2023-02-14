package usecase

import (
	"context"
	"courses-api/model"
	"courses-api/request"
)

type courseCategoryUsecase struct {
	courseCategoryRepository model.CourseCategoryRepository
}

func NewCourseCategoryUsecase(courseCategoryRepository model.CourseCategoryRepository) model.CourseCategoryUsecase {
	return &courseCategoryUsecase{courseCategoryRepository: courseCategoryRepository}
}

func (c *courseCategoryUsecase) StoreCourseCategory(ctx context.Context, req *request.CourseCategoryRequest) (*model.CourseCategory, error) {
	newCourseCategory := &model.CourseCategory{
		Name: req.Name,
	}
	courseCategory, err := c.courseCategoryRepository.Create(ctx, newCourseCategory)
	if err != nil {
		return nil, err
	}
	return courseCategory, nil
}

func (c *courseCategoryUsecase) FetchCourseCategory(ctx context.Context, limit, offset int) ([]*model.CourseCategory, error) {
	courseCategories, err := c.courseCategoryRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return courseCategories, nil
}

func (c *courseCategoryUsecase) GetByID(ctx context.Context, id int64) (*model.CourseCategory, error) {
	courseCategory, err := c.courseCategoryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return courseCategory, nil
}

func (c *courseCategoryUsecase) EditCourseCategory(ctx context.Context, id int64, req *request.CourseCategoryRequest) (*model.CourseCategory, error) {
	_, err := c.courseCategoryRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	courseCategory, err := c.courseCategoryRepository.UpdateByID(ctx, id, &model.CourseCategory{
		Name: req.Name,
	})
	if err != nil {
		return nil, err
	}
	return courseCategory, nil
}

func (c *courseCategoryUsecase) DestroyCourseCategory(ctx context.Context, id int64) error {
	err := c.courseCategoryRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
