package service

import (
	"context"
	"go-micro-sample/user/pkg/domain/entity"
	"go-micro-sample/user/pkg/domain/repository"
)

type IUserService interface {
	GetUser(ctx context.Context, id int64) (*entity.User, error)
	GetUsers(ctx context.Context, ids []int64) ([]*entity.User, error)
	CreateUser(ctx context.Context, name string) (*entity.User, error)
}

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(ur repository.IUserRepository) IUserService {
	return &userService{
		repo: ur,
	}
}

func (us *userService) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	return us.repo.Select(ctx, id)
}

func (us *userService) GetUsers(ctx context.Context, ids []int64) ([]*entity.User, error) {
	return us.repo.SelectAll(ctx, ids)
}

func (us *userService) CreateUser(ctx context.Context, name string) (*entity.User, error) {
	return us.repo.Insert(ctx, name)
}
