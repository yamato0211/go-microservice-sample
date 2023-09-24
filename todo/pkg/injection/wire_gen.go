// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"go-micro-sample/todo/pkg/domain/service"
	"go-micro-sample/todo/pkg/infra"
	"go-micro-sample/todo/pkg/infra/postgres"
	"go-micro-sample/todo/pkg/lib/config"
	"go-micro-sample/todo/pkg/usecase"
)

// Injectors from wire.go:

func InitializeTodoUsecase() usecase.ITodoUsecase {
	dbConfig := config.NewDBConfig()
	db := infra.NewPostgresConnector(dbConfig)
	iTodoRepository := postgres.NewTodoRepository(db)
	iTodoService := service.NewTodoService(iTodoRepository)
	iTodoUsecase := usecase.NewTodoUsecase(iTodoService)
	return iTodoUsecase
}