package module

import (
	usersHTTP "sleekflow/service/users/delivery"
	usersRepository "sleekflow/service/users/repository"
	usersUsecase "sleekflow/service/users/usecase"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		usersHTTP.New,
		usersUsecase.New,
		usersRepository.New,
	),
)
