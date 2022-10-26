package module

import (
	toDoHTTP "sleekflow/service/to_do/delivery"
	toDoRepository "sleekflow/service/to_do/repository"
	toDoUsecase "sleekflow/service/to_do/usecase"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		toDoHTTP.New,
		toDoUsecase.New,
		toDoRepository.New,
	),
)
