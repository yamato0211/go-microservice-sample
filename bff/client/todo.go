package client

import (
	"context"
	"go-micro-sample/bff/graph/model"
	"go-micro-sample/bff/lib/config"
	todopb "go-micro-sample/todo/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ITodoService interface {
	GetTodos(ctx context.Context, id int64) ([]*model.Todo, error)
	CreateTodo(ctx context.Context, title string, id int64) (*model.Todo, error)
	UpdateTodo(ctx context.Context, id int64) (bool, error)
}

type TodoService struct {
	client todopb.TodoServiceClient
}

func NewTodoService(cfg *config.TodoConfig) ITodoService {
	conn, err := grpc.Dial(
		cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := todopb.NewTodoServiceClient(conn)
	return &TodoService{
		client: client,
	}
}

func (ts *TodoService) GetTodos(ctx context.Context, id int64) ([]*model.Todo, error) {
	req := &todopb.GetTodosRequest{
		Id: id,
	}
	res, err := ts.client.GetTodos(ctx, req)
	if err != nil {
		return nil, err
	}
	todos := make([]*model.Todo, len(res.Todos))
	for i, t := range res.Todos {
		todos[i] = &model.Todo{
			ID:     t.GetId(),
			Title:  t.GetTitle(),
			UserID: &t.UserId,
		}
	}
	return todos, nil
}

func (ts *TodoService) CreateTodo(ctx context.Context, title string, id int64) (*model.Todo, error) {
	req := &todopb.CreateTodoRequest{
		Title: title,
		Id:    id,
	}
	todo, err := ts.client.CreateTodo(ctx, req)
	if err != nil {
		return nil, err
	}
	return &model.Todo{
		ID:     todo.GetId(),
		Title:  todo.GetTitle(),
		UserID: &todo.UserId,
	}, nil
}

func (ts *TodoService) UpdateTodo(ctx context.Context, id int64) (bool, error) {
	req := &todopb.UpdateTodoRequest{
		Id: id,
	}
	res, err := ts.client.UpdateTodo(ctx, req)
	if err != nil {
		return res.GetOk(), err
	}
	return res.GetOk(), nil
}
