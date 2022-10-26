package usecase

import (
	"sleekflow/lib/database_transaction"
	"sleekflow/service/to_do"
)

type Usecase struct {
	transactionManager database_transaction.Client
	toDoRepo           to_do.Repository
}

func New(
	toDoRepo to_do.Repository,
	transactionManager database_transaction.Client,
) to_do.Usecase {
	return &Usecase{
		transactionManager: transactionManager,
		toDoRepo:           toDoRepo,
	}
}
