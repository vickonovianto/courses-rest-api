package usecase

import (
	"context"
	"courses-api/model"
	"courses-api/request"
	"errors"

	"gorm.io/gorm"
)

type courseUsecase struct {
	courseRepository         model.CourseRepository
	courseCategoryRepository model.CourseCategoryRepository
}

func NewCourseUsecase(
	courseRepository model.CourseRepository,
	courseCategoryRepository model.CourseCategoryRepository) model.CourseUsecase {
	return &courseUsecase{
		courseRepository:         courseRepository,
		courseCategoryRepository: courseCategoryRepository,
	}
}

func (c *courseUsecase) StoreCourse(ctx context.Context, req *request.CourseRequest) (*model.Course, error) {
	newCourse := &model.Course{
		Title:            req.Title,
		CourseCategoryID: req.CourseCategoryID,
	}
	_, err := c.courseCategoryRepository.FindByID(ctx, req.CourseCategoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("course category id not valid ")
		}
		return nil, err
	}
	course, err := c.courseRepository.Create(ctx, newCourse)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (c *courseUsecase) FetchCourse(ctx context.Context, limit, offset int) ([]*model.Course, error) {
	courses, err := c.courseRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *courseUsecase) GetByID(ctx context.Context, id int64) (*model.Course, error) {
	course, err := c.courseRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (c *courseUsecase) EditCourse(ctx context.Context, id int64, req *request.CourseRequest) (*model.Course, error) {
	_, err := c.courseRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	_, err = c.courseCategoryRepository.FindByID(ctx, req.CourseCategoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("course category id not valid ")
		}
		return nil, err
	}
	course, err := c.courseRepository.UpdateByID(ctx, id, &model.Course{
		Title:            req.Title,
		CourseCategoryID: req.CourseCategoryID,
	})
	if err != nil {
		return nil, err
	}
	return course, nil
}

func (c *courseUsecase) DestroyCourse(ctx context.Context, id int64) error {
	err := c.courseRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
