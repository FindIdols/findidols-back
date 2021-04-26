package entity

import (
	"regexp"
	"strings"
	"time"
)

//Idol data
type BankAccount struct {
	ID        ID
	BankName  string
	Type      string
	Agency    string
	Operation string
	Account   string
	Digit     string
	CreatedAt time.Time
}

//NewIdol create a new idol
func NewBankAccount(bankName, typeAccount, agency, operation, account, digit string) (*BankAccount, error) {
	ba := &BankAccount{
		ID:        NewID(),
		BankName:  strings.TrimSpace(bankName),
		Type:      typeAccount,
		Agency:    strings.TrimSpace(agency),
		Operation: strings.TrimSpace(operation),
		Account:   strings.TrimSpace(account),
		Digit:     digit,
		CreatedAt: time.Now(),
	}

	err := ba.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return ba, nil
}

//Validate validate data
func (ba *BankAccount) Validate() error {

	rgx, _ := regexp.Compile("^[a-zA-Z]*$")

	if ba.BankName == "" || !rgx.MatchString(ba.BankName) {
		return ErrInvalidEntity
	}

	if ba.Type != "corrente" && ba.Type != "poupanca" {
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("^[0-9]*$")

	if ba.Agency == "" || !rgx.MatchString(ba.Agency) {
		return ErrInvalidEntity
	}

	if ba.Account == "" || !rgx.MatchString(ba.Account) {
		return ErrInvalidEntity
	}

	if ba.Digit == "" || !rgx.MatchString(ba.Digit) {
		return ErrInvalidEntity
	}

	return nil
}
