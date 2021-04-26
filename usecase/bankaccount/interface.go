package bankaccount

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Reader interface
type Reader interface {
}

//Writer idol writer
type Writer interface {
	Create(e *entity.BankAccount) (entity.ID, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	CreateBankAccount(
		bankName string,
		typeAccount string,
		agency string,
		operation string,
		account string,
		digit string,
	) (entity.ID, error)
}
