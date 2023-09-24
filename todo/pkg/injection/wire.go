//go:build wireinject
// +build wireinject

package injection

import (
	"go-micro-sample/todo/pkg/domain/service"
	"go-micro-sample/todo/pkg/infra"
	"go-micro-sample/todo/pkg/infra/postgres"
	"go-micro-sample/todo/pkg/lib/config"
	"go-micro-sample/todo/pkg/usecase"

	"github.com/google/wire"
)

func InitializeTodoUsecase() usecase.ITodoUsecase {
	wire.Build(
		config.NewDBConfig,
		infra.NewPostgresConnector,
		postgres.NewTodoRepository,
		service.NewTodoService,
		usecase.NewTodoUsecase,
	)
	return &usecase.TodoUsecase{}
}
