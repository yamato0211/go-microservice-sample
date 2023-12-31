// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package injection

import (
	"go-micro-sample/user/pkg/domain/service"
	"go-micro-sample/user/pkg/infra"
	"go-micro-sample/user/pkg/infra/postgres"
	"go-micro-sample/user/pkg/lib/config"
	"go-micro-sample/user/pkg/usecase"
)

// Injectors from wire.go:

func InitializeUserUsecase() usecase.IUserUsecase {
	dbConfig := config.NewDBConfig()
	db := infra.NewPostgresConnector(dbConfig)
	iUserRepository := postgres.NewUserRepository(db)
	iUserService := service.NewUserService(iUserRepository)
	iUserUsecase := usecase.NewUserUsecase(iUserService)
	return iUserUsecase
}
