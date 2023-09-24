package client

import (
	"context"
	"go-micro-sample/bff/graph/model"
	"go-micro-sample/bff/lib/config"
	userpb "go-micro-sample/user/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IUserService interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)
	CreateUser(ctx context.Context, name string) (*model.User, error)
}

type UserService struct {
	client userpb.UserServiceClient
}

func NewUserService(cfg *config.UserConfig) IUserService {
	conn, err := grpc.Dial(
		cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal(err)
	}
	client := userpb.NewUserServiceClient(conn)
	return &UserService{
		client: client,
	}
}

func (us *UserService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	req := &userpb.GetUserRequest{
		Id: id,
	}
	user, err := us.client.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   int(user.GetId()),
		Name: user.GetName(),
	}, nil
}

func (us *UserService) CreateUser(ctx context.Context, name string) (*model.User, error) {
	req := &userpb.CreateUserRequest{
		Name: name,
	}
	user, err := us.client.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   int(user.GetId()),
		Name: user.GetName(),
	}, nil
}
