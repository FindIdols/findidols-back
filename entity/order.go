package entity

import (
	"time"
)

//Order data
type Order struct {
	ID          ID
	OrderNumber int
	Content
	IdolID     string
	UserID     string
	TermsOfUse bool
	CreatedAt  time.Time
}

//NewOrder create a new order
func NewOrder(user User, content Content, termsOfUse bool, idolID string) (*Order, error) {
	o := &Order{
		ID:         NewID(),
		Content:    content,
		TermsOfUse: termsOfUse,
		IdolID:     idolID,
		UserID:     user.ID.String(),
		CreatedAt:  time.Now(),
	}

	err := o.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return o, nil
}

//Validate validate data
func (o *Order) Validate() error {

	if o.IdolID == "" {
		return ErrInvalidEntity
	}

	if o.TermsOfUse == false {
		return ErrInvalidEntity
	}

	return nil
}
