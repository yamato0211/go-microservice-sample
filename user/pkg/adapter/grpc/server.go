package grpc

import (
	"context"
	"go-micro-sample/user/pkg/usecase"
	userpb "go-micro-sample/user/proto"
)

type userServer struct {
	userpb.UnimplementedUserServiceServer
	usecase usecase.IUserUsecase
}

func NewUserServer(usecase usecase.IUserUsecase) *userServer {
	return &userServer{
		usecase: usecase,
	}
}

func (s *userServer) GetUser(ctx context.Context, gur *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := s.usecase.GetUser(ctx, gur.GetId())
	if err != nil {
		return nil, err
	}
	return &userpb.GetUserResponse{
		Id:   user.ID,
		Name: user.Name,
	}, nil
}

func (s *userServer) GetUsers(ctx context.Context, gur *userpb.GetUsersRequest) (*userpb.GetUsersResponse, error) {
	users, err := s.usecase.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	res := make([]*userpb.GetUserResponse, len(users))
	for i, u := range users {
		res[i] = &userpb.GetUserResponse{
			Id:   u.ID,
			Name: u.Name,
		}
	}
	return &userpb.GetUsersResponse{
		Users: res,
	}, nil
}

func (s *userServer) CreateUser(ctx context.Context, cur *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	user, err := s.usecase.CreateUser(ctx, cur.GetName())
	if err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{
		Id:   user.ID,
		Name: user.Name,
	}, nil
}
