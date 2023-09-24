package usecase

import (
	"context"
	"go-micro-sample/todo/pkg/domain/entity"
	"go-micro-sample/todo/pkg/domain/service"
)

type ITodoUsecase interface {
	GetTodos(ctx context.Context, id int64) ([]*entity.Todo, error)
	CreateTodo(ctx context.Context, title string, id int64) (*entity.Todo, error)
	UpdateTodo(ctx context.Context, id int64) (bool, error)
}

type TodoUsecase struct {
	svc service.ITodoService
}

func NewTodoUsecase(ts service.ITodoService) ITodoUsecase {
	return &TodoUsecase{
		svc: ts,
	}
}

func (tu *TodoUsecase) GetTodos(ctx context.Context, id int64) ([]*entity.Todo, error) {
	todos, err := tu.svc.GetTodos(ctx, id)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tu *TodoUsecase) CreateTodo(ctx context.Context, title string, id int64) (*entity.Todo, error) {
	todo, err := tu.svc.CreateTodo(ctx, title, id)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func (tu *TodoUsecase) UpdateTodo(ctx context.Context, id int64) (bool, error) {
	return tu.svc.UpdateTodo(ctx, id)
}
