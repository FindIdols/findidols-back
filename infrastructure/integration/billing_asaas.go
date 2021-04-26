package integration

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/FindIdols/findidols-back/entity"
)

type CreditCard struct {
	HolderName  string `json:"holderName"`
	Number      string `json:"number"`
	ExpiryMonth string `json:"expiryMonth"`
	ExpiryYear  string `json:"expiryYear"`
	CCV         string `json:"ccv"`
}

type CreditCardHolderInfo struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	CpfCnpj           string `json:"cpfCnpj"`
	PostalCode        string `json:"postalCode"`
	AddressNumber     string `json:"addressNumber"`
	AddressComplement string `json:"addressComplement"`
}

func CreateBilling(billing entity.Billing) ([]byte, error) {

	paymentJSON, err := json.Marshal(map[string]interface{}{
		"customer":    billing.CustomerID,
		"billingType": billing.BillingType,
		"dueDate":     billing.DueDate,
		"value":       billing.Value,
		"description": billing.Description,
	})
	fmt.Println(string(paymentJSON))

	responseBody := bytes.NewBuffer(paymentJSON)

	if err != nil {
		fmt.Println("Erro ao gerar json", err)
		return nil, errors.New("Erro ao gerar json")
	}

	req, err := http.NewRequest("POST", "https://www.asaas.com/api/v3/payments", responseBody)
	req.Header.Set("access_token", "36d24f6cb5871a629a727deb0f652ed20f6be4d9bebb4b28ca50b9749dfc2e13")

	// req, err := http.NewRequest("POST", "https://sandbox.asaas.com/api/v3/payments", responseBody)
	// req.Header.Set("access_token", "7fde62b9f2a45dc9fd40a5b378869e14f4a4477cce8de0bc687cfa6b55f3ac62")

	if err != nil {
		fmt.Println("Erro ao fazer body", err)
		return nil, errors.New("Erro ao fazer body")
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Erro ao efetuar request", err)
		return nil, errors.New("Erro ao efetuar request")
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Erro ao ler body", string(data))
		return nil, errors.New("Erro ao ler body")
	}

	if res.StatusCode == 400 {
		fmt.Println("Erro ao efetuar request com codigo 400", string(data))
		return nil, errors.New("Erro ao efetuar request com codigo 400")
	}

	if res.StatusCode == 500 {
		fmt.Println("Erro ao efetuar request com codigo 500", string(data))
		return nil, errors.New("Erro ao efetuar request com codigo 500")
	}

	return data, nil
}

func CreateCreditCardBilling(billing entity.Billing) ([]byte, error) {

	paymentJSON, err := json.Marshal(map[string]interface{}{
		"customer":    billing.CustomerID,
		"billingType": billing.BillingType,
		"dueDate":     billing.DueDate,
		"value":       billing.Value,
		"description": billing.Description,
		"creditCard": &CreditCard{
			billing.CreditCardName,
			billing.CreditCardNumber,
			billing.CreditCardExpiryMonth,
			billing.CreditCardExpiryYear,
			billing.CCV,
		},
		"creditCardHolderInfo": &CreditCardHolderInfo{
			billing.CreditCardNameHolder,
			billing.CreditCardEmailHolder,
			billing.CreditCardPhoneHolder,
			billing.Customer.CPF,
			billing.Customer.CEP,
			billing.Customer.AddresNumber,
			billing.Customer.AddressComplment,
		},
	})

	fmt.Println(string(paymentJSON))

	responseBody := bytes.NewBuffer(paymentJSON)

	if err != nil {
		fmt.Println("Erro ao gerar json", err)
		return nil, errors.New("Erro ao gerar json")
	}

	req, err := http.NewRequest("POST", "https://www.asaas.com/api/v3/payments", responseBody)
	req.Header.Set("access_token", "36d24f6cb5871a629a727deb0f652ed20f6be4d9bebb4b28ca50b9749dfc2e13")

	// req, err := http.NewRequest("POST", "https://sandbox.asaas.com/api/v3/payments", responseBody)
	// req.Header.Set("access_token", "7fde62b9f2a45dc9fd40a5b378869e14f4a4477cce8de0bc687cfa6b55f3ac62")

	if err != nil {
		fmt.Println("Erro ao fazer body", err)
		return nil, errors.New("Erro ao fazer body")
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		fmt.Println("Erro ao efetuar request", err)
		return nil, errors.New("Erro ao efetuar request")
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println("Erro ao ler body", string(data))
		return nil, errors.New("Erro ao ler body")
	}

	if res.StatusCode == 400 {
		fmt.Println("Erro ao efetuar request com codigo 400", string(data))
		return nil, errors.New("Erro ao efetuar request com codigo 400")
	}

	if res.StatusCode == 500 {
		fmt.Println("Erro ao efetuar request com codigo 500", string(data))
		return nil, errors.New("Erro ao efetuar request com codigo 500")
	}

	return data, nil
}
