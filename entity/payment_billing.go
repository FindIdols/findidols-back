package entity

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

//Content data
type Billing struct {
	Customer
	CustomerID            string
	BillingType           string
	Value                 float64
	DueDate               string
	Description           string
	InstallmentCount      int
	InstallmentValue      float64
	CreditCardNumber      string
	CreditCardName        string
	CreditCardNameHolder  string
	CreditCardPhoneHolder string
	CreditCardEmailHolder string
	CreditCardExpiryMonth string
	CreditCardExpiryYear  string
	CCV                   string
}

//NewCustomer create a new customer
func NewBilling(
	customer Customer,
	customerID string,
	description string,
	value,
	installmentValue float64,
	installmentCount int,
	billingType string,
	creditCardNumber string,
	creditCardName string,
	creditCardNameHolder string,
	creditCardPhoneHolder string,
	creditCardEmailHolder string,
	creditCardExpiryMonth string,
	creditCardExpiryYear string,
	ccv string,
) (*Billing, error) {
	b := &Billing{
		Customer:              customer,
		CustomerID:            customerID,
		DueDate:               time.Now().AddDate(0, 0, 3).Format("2006-01-02"),
		Description:           description,
		Value:                 value,
		InstallmentValue:      installmentValue,
		BillingType:           billingType,
		InstallmentCount:      installmentCount,
		CreditCardNumber:      strings.TrimSpace(creditCardNumber),
		CreditCardName:        strings.TrimSpace(creditCardName),
		CreditCardNameHolder:  strings.TrimSpace(creditCardNameHolder),
		CreditCardPhoneHolder: strings.TrimSpace(creditCardPhoneHolder),
		CreditCardEmailHolder: strings.TrimSpace(creditCardEmailHolder),
		CreditCardExpiryMonth: creditCardExpiryMonth,
		CreditCardExpiryYear:  creditCardExpiryYear,
		CCV:                   strings.TrimSpace(ccv),
	}

	if billingType == "BOLETO" {
		err := b.ValidateBankSlipPayment()

		if err != nil {
			return nil, ErrInvalidEntity
		}

		return b, nil
	}

	err := b.ValidateCreditCardPayment()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return b, nil
}

// //Validate validate data
func (b *Billing) ValidateBankSlipPayment() error {

	if b.CustomerID == "" {
		return ErrInvalidEntity
	}

	if b.Description == "" {
		return ErrInvalidEntity
	}

	if b.Value == 0 && b.Value < 0 {
		return ErrInvalidEntity
	}

	if b.BillingType != "BOLETO" {
		return ErrInvalidEntity
	}

	return nil
}

func (b *Billing) ValidateCreditCardPayment() error {

	if b.CustomerID == "" {
		fmt.Println("id do cliente invalido")
		return ErrInvalidEntity
	}

	if b.Description == "" {
		fmt.Println("descricao do pedido invalido")
		return ErrInvalidEntity
	}

	if b.Value == 0 || b.Value < 0 {
		fmt.Println("valor do pedido invalido")
		return ErrInvalidEntity
	}

	if b.BillingType != "CREDIT_CARD" {
		return ErrInvalidEntity
	}

	rgx, _ := regexp.Compile("^[A-Za-z ]*$")

	if b.CreditCardName == "" || !rgx.MatchString(b.CreditCardName) {
		fmt.Println("nome do cartao invalido")
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("^[A-Za-z \u00C0-\u00FF]*$")

	if b.CreditCardNameHolder == "" || !rgx.MatchString(b.CreditCardNameHolder) {
		fmt.Println("nome do dono do cartao invalido")
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("^[0-9]*$")

	if b.CreditCardNumber == "" || !rgx.MatchString(b.CreditCardNumber) {
		fmt.Println("numero do cartao invalido")
		return ErrInvalidEntity
	}

	if b.CreditCardPhoneHolder == "" || !rgx.MatchString(b.CreditCardPhoneHolder) || len(b.CreditCardPhoneHolder) != 11 {
		fmt.Println("celular do dono do cartao invalido")
		return ErrInvalidEntity
	}

	if b.CCV == "" || !rgx.MatchString(b.CCV) {
		fmt.Println("cvv do cartao invalido")
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("^[A-Za-z0-9@.]*$")

	if b.CreditCardEmailHolder == "" || !rgx.MatchString(b.CreditCardEmailHolder) {
		fmt.Println("email do dono do cartao invalido")
		return ErrInvalidEntity
	}

	if b.CreditCardExpiryMonth == "" {
		return ErrInvalidEntity
	}

	if b.CreditCardExpiryYear == "" {
		return ErrInvalidEntity
	}

	return nil
}
