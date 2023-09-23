package usecase

import (
	"context"
	"go-micro-sample/user/pkg/domain/entity"
	"go-micro-sample/user/pkg/domain/service"
)

type IUserUsecase interface {
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, name string) (*entity.User, error)
}

type UserUsecase struct {
	svc service.IUserService
}

func NewUserUsecase(us service.IUserService) IUserUsecase {
	return &UserUsecase{
		svc: us,
	}
}

func (uu *UserUsecase) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	user, err := uu.svc.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu *UserUsecase) GetUsers(ctx context.Context) ([]*entity.User, error) {
	users, err := uu.svc.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu *UserUsecase) CreateUser(ctx context.Context, name string) (*entity.User, error) {
	user, err := uu.svc.CreateUser(ctx, name)
	if err != nil {
		return nil, err
	}
	return user, nil
}
