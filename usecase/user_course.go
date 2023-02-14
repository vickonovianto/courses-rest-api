package usecase

import (
	"context"
	"courses-api/model"
	"courses-api/request"
	"errors"

	"gorm.io/gorm"
)

type userCourseUsecase struct {
	userCourseRepository model.UserCourseRepository
	userRepository       model.UserRepository
	courseRepository     model.CourseRepository
}

func NewUserCourseUsecase(
	userCourseRepository model.UserCourseRepository,
	userRepository model.UserRepository,
	courseRepository model.CourseRepository) model.UserCourseUsecase {
	return &userCourseUsecase{
		userCourseRepository: userCourseRepository,
		userRepository:       userRepository,
		courseRepository:     courseRepository,
	}
}

func (u *userCourseUsecase) StoreUserCourse(ctx context.Context, req *request.UserCourseRequest) (*model.UserCourse, error) {
	newUserCourse := &model.UserCourse{
		UserID:   req.UserID,
		CourseID: req.CourseID,
	}
	_, err := u.userRepository.FindByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user id not valid ")
		}
		return nil, err
	}
	_, err = u.courseRepository.FindByID(ctx, req.CourseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("course id not valid ")
		}
		return nil, err
	}
	userCourse, err := u.userCourseRepository.Create(ctx, newUserCourse)
	if err != nil {
		return nil, err
	}
	return userCourse, nil
}

func (u *userCourseUsecase) FetchUserCourse(ctx context.Context, limit, offset int) ([]*model.UserCourse, error) {
	userCourses, err := u.userCourseRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return userCourses, nil
}

func (u *userCourseUsecase) GetByID(ctx context.Context, id int64) (*model.UserCourse, error) {
	userCourse, err := u.userCourseRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return userCourse, nil
}

func (u *userCourseUsecase) EditUserCourse(ctx context.Context, id int64, req *request.UserCourseRequest) (*model.UserCourse, error) {
	_, err := u.userCourseRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	_, err = u.userRepository.FindByID(ctx, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user id not valid ")
		}
		return nil, err
	}
	_, err = u.courseRepository.FindByID(ctx, req.CourseID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("course id not valid ")
		}
		return nil, err
	}
	userCourse, err := u.userCourseRepository.UpdateByID(ctx, id, &model.UserCourse{
		UserID:   req.UserID,
		CourseID: req.CourseID,
	})
	if err != nil {
		return nil, err
	}
	return userCourse, nil
}

func (u *userCourseUsecase) DestroyUserCourse(ctx context.Context, id int64) error {
	err := u.userCourseRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
