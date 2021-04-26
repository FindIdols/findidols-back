package entity

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

//User data
type User struct {
	ID        ID
	FirstName string
	LastName  string
	Email     string
	Phone     string
	Genre     string
	Category  string
	CreatedAt time.Time
}

//NewUser create a new user
func NewUser(firstName, lastName, email, phone, genre, category string) (*User, error) {
	u := &User{
		ID:        NewID(),
		Email:     strings.TrimSpace(email),
		FirstName: strings.TrimSpace(firstName),
		LastName:  strings.TrimSpace(lastName),
		Phone:     strings.TrimSpace(phone),
		Genre:     strings.TrimSpace(genre),
		Category:  category,
		CreatedAt: time.Now(),
	}

	err := u.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return u, nil
}

//Validate validate data
func (u *User) Validate() error {
	rgx, _ := regexp.Compile("^[A-Za-z0-9@.]*$")

	if u.Email == "" || !rgx.MatchString(u.Email) {
		fmt.Println("email_invalido")
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("^[A-Za-z \u00C0-\u00FF]*$")

	if u.FirstName == "" || !rgx.MatchString(u.FirstName) {
		fmt.Println("nome_invalido")
		return ErrInvalidEntity
	}

	if u.LastName == "" || !rgx.MatchString(u.LastName) {
		fmt.Println("sobrenome_invalido")
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("^[0-9]*$")

	if u.Phone == "" || !rgx.MatchString(u.Phone) || len(u.Phone) != 11 {
		fmt.Println("telefone_invalido")
		return ErrInvalidEntity
	}

	if u.Genre == "" || u.Genre != "male" && u.Genre != "female" {
		return ErrInvalidEntity
	}

	if u.Category == "" || u.Category != "fan" && u.Category != "idol" {
		return ErrInvalidEntity
	}

	return nil
}
