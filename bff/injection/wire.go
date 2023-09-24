//go:build wireinject
// +build wireinject

package injection

import (
	"go-micro-sample/bff/client"
	"go-micro-sample/bff/lib/config"

	"github.com/google/wire"
)

func InitializeTodoService() client.ITodoService {
	wire.Build(
		config.NewTodoConfig,
		client.NewTodoService,
	)
	return &client.TodoService{}
}

func InitializeUserService() client.IUserService {
	wire.Build(
		config.NewUserConfig,
		client.NewUserService,
	)
	return &client.UserService{}
}
