package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/FindIdols/findidols-back/api/presenter"
	"github.com/FindIdols/findidols-back/usecase/billing"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func createCustomer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error_create_customer"

		var input struct {
			Name             string `json:"name"`
			Email            string `json:"email"`
			Phone            string `json:"phone"`
			CPF              string `json:"cpf"`
			CEP              string `json:"cep"`
			AddressName      string `json:"addressName"`
			AddresNumber     string `json:"addressNumber"`
			AddressComplment string `json:"addressComplment"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			fmt.Println("erro decode")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := billing.CreateCustomer(
			input.Name,
			input.Email,
			input.Phone,
			input.CPF,
			input.CEP,
			input.AddressName,
			input.AddresNumber,
			input.AddressComplment,
		)

		var customerInfo presenter.Customer

		json.Unmarshal(data, &customerInfo)

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(customerInfo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

func createBilling() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		errorMessage := "error_create_billing"

		var input struct {
			Name                  string  `json:"name"`
			Email                 string  `json:"email"`
			Phone                 string  `json:"phone"`
			CPF                   string  `json:"cpf"`
			CEP                   string  `json:"cep"`
			AddressName           string  `json:"addressName"`
			AddresNumber          string  `json:"addressNumber"`
			AddressComplment      string  `json:"addressComplment"`
			BillingType           string  `json:"billingType"`
			Value                 float64 `json:"value"`
			Description           string  `json:"description"`
			InstallmentCount      int     `json:"installmentCount"`
			InstallmentValue      float64 `json:"installmentValue"`
			CreditCardNameHolder  string  `json:"creditCardNameHolder"`
			CreditCardPhoneHolder string  `json:"creditCardPhoneHolder"`
			CreditCardEmailHolder string  `json:"creditCardEmailHolder"`
			CreditCardName        string  `json:"creditCardName"`
			CreditCardNumber      string  `json:"creditCardNumber"`
			CreditCardExpiryMonth string  `json:"creditCardExpiryMonth"`
			CreditCardExpiryYear  string  `json:"creditCardExpiryYear"`
			CCV                   string  `json:"ccv"`
		}

		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			fmt.Println("erro decode")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		data, err := billing.CreateCustomer(
			input.Name,
			input.Email,
			input.Phone,
			input.CPF,
			input.CEP,
			input.AddressName,
			input.AddresNumber,
			input.AddressComplment,
		)

		if err != nil {
			fmt.Println("erro_criacao_customer")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		var customerInfo presenter.Customer

		json.Unmarshal(data, &customerInfo)

		fmt.Println(customerInfo)

		data, err = billing.CreateBilling(
			customerInfo.ID,
			input.Name,
			input.Email,
			input.Phone,
			input.CPF,
			input.CEP,
			input.AddressName,
			input.AddresNumber,
			input.AddressComplment,
			input.BillingType,
			input.Value,
			input.Description,
			input.InstallmentCount,
			input.InstallmentValue,
			input.CreditCardNumber,
			input.CreditCardName,
			input.CreditCardNameHolder,
			input.CreditCardEmailHolder,
			input.CreditCardPhoneHolder,
			input.CreditCardExpiryMonth,
			input.CreditCardExpiryYear,
			input.CCV,
		)

		if err != nil {
			fmt.Println("erro_criacao_pagamento")
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}

		var billingInfo presenter.Billing

		json.Unmarshal(data, &billingInfo)

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(billingInfo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errorMessage))
			return
		}
	})
}

//MakeBillingHandlers make url handlers
func MakeBillingHandlers(r *mux.Router, n negroni.Negroni) {
	r.Handle("/v1/billing/customer", n.With(
		negroni.Wrap(createCustomer()),
	)).Methods("POST", "OPTIONS").Name("createCustomer")

	r.Handle("/v1/billing/payment", n.With(
		negroni.Wrap(createBilling()),
	)).Methods("POST", "OPTIONS").Name("createBilling")

}
