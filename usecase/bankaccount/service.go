package bankaccount

import (
	"fmt"

	"github.com/FindIdols/findidols-back/entity"
)

//Service book usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//CreateBankAccount create a bank account
func (s *Service) CreateBankAccount(
	bankName string,
	typeAccount string,
	agency string,
	operation string,
	account string,
	digit string,
) (entity.ID, error) {

	bankaccount, err := entity.NewBankAccount(bankName, typeAccount, agency, operation, account, digit)

	if err != nil {
		fmt.Println("Erro ao criar conta bancária")
		return bankaccount.ID, err
	}

	return s.repo.Create(bankaccount)
}
