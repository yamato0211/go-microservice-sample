package service

import (
	"context"
	"go-micro-sample/todo/pkg/domain/entity"
	"go-micro-sample/todo/pkg/domain/repository"
)

type ITodoService interface {
	GetTodos(ctx context.Context, id int64) ([]*entity.Todo, error)
	CreateTodo(ctx context.Context, title string, id int64) (*entity.Todo, error)
	UpdateTodo(ctx context.Context, id int64) (bool, error)
}

type todoService struct {
	repo repository.ITodoRepository
}

func NewTodoService(tr repository.ITodoRepository) ITodoService {
	return &todoService{
		repo: tr,
	}
}

func (ts *todoService) GetTodos(ctx context.Context, id int64) ([]*entity.Todo, error) {
	return ts.repo.SelectAll(ctx, id)
}

func (ts *todoService) CreateTodo(ctx context.Context, title string, id int64) (*entity.Todo, error) {
	return ts.repo.Insert(ctx, title, id)
}

func (ts *todoService) UpdateTodo(ctx context.Context, id int64) (bool, error) {
	return ts.repo.Update(ctx, id)
}
