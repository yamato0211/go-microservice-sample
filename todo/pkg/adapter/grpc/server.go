package grpc

import (
	"context"
	"go-micro-sample/todo/pkg/usecase"
	todopb "go-micro-sample/todo/proto"
)

type todoServer struct {
	todopb.UnimplementedTodoServiceServer
	usecase usecase.ITodoUsecase
}

func NewTodoServer(usecase usecase.ITodoUsecase) *todoServer {
	return &todoServer{
		usecase: usecase,
	}
}

func (s *todoServer) GetTodos(ctx context.Context, gtr *todopb.GetTodosRequest) (*todopb.GetTodosResponse, error) {
	todos, err := s.usecase.GetTodos(ctx, gtr.GetId())
	if err != nil {
		return nil, err
	}
	res := make([]*todopb.GetTodoResponse, len(todos))
	for i, t := range todos {
		res[i] = &todopb.GetTodoResponse{
			Id:     t.ID,
			Title:  t.Title,
			Done:   t.Done,
			UserId: t.UserID,
		}
	}
	return &todopb.GetTodosResponse{
		Todos: res,
	}, nil
}

func (s *todoServer) CreateTodo(ctx context.Context, ctr *todopb.CreateTodoRequest) (*todopb.CreateTodoResponse, error) {
	todo, err := s.usecase.CreateTodo(ctx, ctr.GetTitle(), ctr.GetId())
	if err != nil {
		return nil, err
	}
	return &todopb.CreateTodoResponse{
		Id:     todo.ID,
		Title:  todo.Title,
		UserId: todo.UserID,
	}, nil
}

func (s *todoServer) UpdateTodo(ctx context.Context, utr *todopb.UpdateTodoRequest) (*todopb.UpdateTodoResponse, error) {
	ok, err := s.usecase.UpdateTodo(ctx, utr.GetId())
	if err != nil {
		return nil, err
	}
	return &todopb.UpdateTodoResponse{
		Ok: ok,
	}, nil
}
