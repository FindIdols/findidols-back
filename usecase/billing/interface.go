package billing

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Reader interface
type Reader interface {
}

//Writer customer writer
type Writer interface {
	CreateCustomer(
		email string,
		phone string,
		name string,
		cpf string,
		cep string,
		addressName string,
		addresNumber string,
		addressComplment string,
	) (entity.ID, error)
	CreateBilling(
		customer entity.Customer,
		customerID string,
		billingType []byte,
		value float64,
		description string,
		installmentCount int,
		installmentValue float64,
		creditCardNumber,
		creditCardName,
		creditCardNameHolder,
		creditCardEmailHolder,
		creditCardPhoneHolder,
		creditCardExpiryMonth,
		creditCardExpiryYear,
		ccv string,
	) (entity.ID, error)
}

//UseCase interface
type UseCase interface {
	CreateCustomer(
		email string,
		phone string,
		name string,
		cpf string,
		cep string,
		addressName string,
		addresNumber string,
		addressComplment string,
	) (entity.ID, error)
	CreateBilling(
		customer entity.Customer,
		customerID string,
		billingType []byte,
		value float64,
		description string,
		installmentCount int,
		installmentValue float64,
		creditCardNumber,
		creditCardName,
		creditCardNameHolder,
		creditCardEmailHolder,
		creditCardPhoneHolder,
		creditCardExpiryMonth,
		creditCardExpiryYear,
		ccv string,
	) (entity.ID, error)
}
