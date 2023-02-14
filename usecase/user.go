package usecase

import (
	"context"
	"courses-api/model"
	"courses-api/request"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository model.UserRepository
}

func NewUserUsecase(userRepository model.UserRepository) model.UserUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (u *userUsecase) StoreUser(ctx context.Context, req *request.UserRequest) (*model.User, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(bytes)
	user, err := u.userRepository.Create(ctx, &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) FetchUser(ctx context.Context, limit, offset int) ([]*model.User, error) {
	users, err := u.userRepository.Fetch(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) EditUser(ctx context.Context, id int64, req *request.UserRequest) (*model.User, error) {
	_, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, err
	}
	hashedPassword := string(bytes)
	user, err := u.userRepository.UpdateByID(ctx, id, &model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) DestroyUser(ctx context.Context, id int64) error {
	err := u.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
