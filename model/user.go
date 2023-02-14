package model

import (
	"context"
	"courses-api/request"
)

type (
	User struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserRepository interface {
		Create(ctx context.Context, user *User) (*User, error)
		Fetch(ctx context.Context, limit, offset int) ([]*User, error)
		FindByID(ctx context.Context, id int64) (*User, error)
		UpdateByID(ctx context.Context, id int64, user *User) (*User, error)
		Delete(ctx context.Context, id int64) error
	}

	UserUsecase interface {
		StoreUser(ctx context.Context, req *request.UserRequest) (*User, error)
		FetchUser(ctx context.Context, limit, offset int) ([]*User, error)
		GetByID(ctx context.Context, id int64) (*User, error)
		EditUser(ctx context.Context, id int64, req *request.UserRequest) (*User, error)
		DestroyUser(ctx context.Context, id int64) error
	}
)

// override gorm table name
func (User) TableName() string {
	return "users"
}
