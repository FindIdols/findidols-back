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

func CreateCustomer(customer entity.Customer) ([]byte, error) {

	customerJSON, err := json.Marshal(map[string]interface{}{
		"name":                 customer.Name,
		"email":                customer.Email,
		"phone":                customer.Phone,
		"mobilePhone":          customer.Phone,
		"cpfCnpj":              customer.CPF,
		"postalCode":           customer.CEP,
		"address":              customer.AddressName,
		"addressNumber":        customer.AddresNumber,
		"complement":           customer.AddressComplment,
		"notificationDisabled": false,
	})

	responseBody := bytes.NewBuffer(customerJSON)

	if err != nil {
		fmt.Println("Erro ao gerar json", err)
		return nil, errors.New("Erro ao gerar json")
	}

	req, err := http.NewRequest("POST", "https://www.asaas.com/api/v3/customers", responseBody)
	req.Header.Set("access_token", "key")

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
