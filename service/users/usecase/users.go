package usecase

import (
	"sleekflow/lib/database_transaction"
	"sleekflow/service/users"
)

type Usecase struct {
	transactionManager database_transaction.Client
	usersRepo          users.Repository
}

func New(
	usersRepo users.Repository,
	transactionManager database_transaction.Client,
) users.Usecase {
	return &Usecase{
		transactionManager: transactionManager,
		usersRepo:          usersRepo,
	}
}
