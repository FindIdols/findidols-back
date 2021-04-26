package billing

import (
	"fmt"

	"github.com/FindIdols/findidols-back/entity"
	"github.com/FindIdols/findidols-back/infrastructure/integration"
)

//CreateCustomer create a customer
func CreateCustomer(
	name string,
	email string,
	phone string,
	cpf string,
	cep string,
	addressName string,
	addresNumber string,
	addressComplment string,
) ([]byte, error) {
	customer, err := entity.NewCustomer(name, email, phone, cpf, cep, addressName, addresNumber, addressComplment)

	if err != nil {
		fmt.Println("cliente_asaas_invalido")
		return nil, err
	}

	data, err := integration.CreateCustomer(*customer)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func CreateBilling(
	customerID string,
	name string,
	email string,
	phone string,
	cpf string,
	cep string,
	addressName string,
	addresNumber string,
	addressComplment string,
	billingType string,
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
) ([]byte, error) {
	customer, err := entity.NewCustomer(name, email, phone, cpf, cep, addressName, addresNumber, addressComplment)

	if err != nil {
		fmt.Println("pagamento_asaas_cliente_invalido")
		fmt.Println(err)
		return nil, err
	}

	billing, err := entity.NewBilling(
		*customer,
		customerID,
		description,
		value,
		installmentValue,
		installmentCount,
		billingType,
		creditCardNumber,
		creditCardName,
		creditCardNameHolder,
		creditCardPhoneHolder,
		creditCardEmailHolder,
		creditCardExpiryMonth,
		creditCardExpiryYear,
		ccv,
	)

	if err != nil {
		fmt.Println("pagamento_asaas_pagamento_invalido")
		fmt.Println(err)
		return nil, err
	}

	if billingType == "BOLETO" {
		data, err := integration.CreateBilling(*billing)

		if err != nil {
			return nil, err
		}

		return data, nil
	}

	fmt.Println("Credito")
	data, err := integration.CreateCreditCardBilling(*billing)

	if err != nil {
		return nil, err
	}

	return data, nil
}
