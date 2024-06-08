package transactionservice

import "clean-code-structure/service/transactionservice/adapter/simpletransaction"

type Transaction interface {
	GetRollbackChannel() <-chan bool
	GetCommitChannel() <-chan bool
	Rollback()
	Commit()
}

type Service struct {
	adapter Transaction
}

func New() Service {
	return Service{}
}

func (s Service) NewTransaction() Transaction {
	return simpletransaction.New()
}
