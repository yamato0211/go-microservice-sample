//go:build wireinject
// +build wireinject

package injection

import (
	"go-micro-sample/user/pkg/domain/service"
	"go-micro-sample/user/pkg/infra"
	"go-micro-sample/user/pkg/infra/postgres"
	"go-micro-sample/user/pkg/lib/config"
	"go-micro-sample/user/pkg/usecase"

	"github.com/google/wire"
)

func InitializeUserUsecase() usecase.IUserUsecase {
	wire.Build(
		config.NewDBConfig,
		infra.NewPostgresConnector,
		postgres.NewUserRepository,
		service.NewUserService,
		usecase.NewUserUsecase,
	)
	return &usecase.UserUsecase{}
}
