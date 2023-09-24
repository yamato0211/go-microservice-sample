package repository

import (
	"context"
	"go-micro-sample/user/pkg/domain/entity"
)

type IUserRepository interface {
	SelectAll(ctx context.Context, ids []int64) ([]*entity.User, error)
	Select(ctx context.Context, id int64) (*entity.User, error)
	Insert(ctx context.Context, name string) (*entity.User, error)
}
