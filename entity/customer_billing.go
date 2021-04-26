package entity

import (
	"fmt"
	"regexp"
	"strings"
)

//Content data
type Customer struct {
	Email            string
	Phone            string
	Name             string
	CPF              string
	CEP              string
	AddressName      string
	AddresNumber     string
	AddressComplment string
}

//NewCustomer create a new customer
func NewCustomer(name, email, phone, cpf, cep, addressName, addressNumber, addressComplement string) (*Customer, error) {
	c := &Customer{
		Name:             strings.TrimSpace(name),
		Email:            strings.TrimSpace(email),
		Phone:            strings.TrimSpace(phone),
		CPF:              strings.TrimSpace(cpf),
		CEP:              strings.TrimSpace(cep),
		AddressName:      strings.TrimSpace(addressName),
		AddresNumber:     strings.TrimSpace(addressNumber),
		AddressComplment: strings.TrimSpace(addressComplement),
	}

	err := c.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return c, nil
}

//Validate validate data
func (c *Customer) Validate() error {

	rgx, _ := regexp.Compile("^[A-Za-z \u00C0-\u00FF]*$")

	if c.Name == "" || !rgx.MatchString(c.Name) {
		fmt.Println("nome do cliente invalido")
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("^[A-Za-z0-9@.]*$")

	if c.Email == "" || !rgx.MatchString(c.Email) {
		fmt.Println("email do cliente invalido")
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("^[0-9]*$")

	if c.Phone == "" || !rgx.MatchString(c.Phone) || len(c.Phone) != 11 {
		fmt.Println("celular do cliente invalido")
		return ErrInvalidEntity
	}

	if c.CPF == "" || !rgx.MatchString(c.CPF) || len(c.CPF) != 11 {
		fmt.Println("cpf do cliente invalido")
		return ErrInvalidEntity
	}

	if c.CEP == "" || !rgx.MatchString(c.CEP) || len(c.CEP) != 8 {
		fmt.Println("cep do cliente invalido")
		return ErrInvalidEntity
	}

	if c.AddresNumber == "" || !rgx.MatchString(c.AddresNumber) {
		fmt.Println("numero do endereco do cliente invalido")
		return ErrInvalidEntity
	}

	if c.AddressName == "" {
		fmt.Println("nome do cliente invalido")
		return ErrInvalidEntity
	}

	return nil
}
