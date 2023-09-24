package repository

import (
	"context"
	"go-micro-sample/todo/pkg/domain/entity"
)

type ITodoRepository interface {
	SelectAll(ctx context.Context, id int64) ([]*entity.Todo, error)
	Insert(ctx context.Context, title string, id int64) (*entity.Todo, error)
	Update(ctx context.Context, id int64) (bool, error)
}
