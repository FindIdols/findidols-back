package entity

import (
	"regexp"
	"strings"
	"time"
)

//Idol data
type Idol struct {
	ID               ID
	ArtisticName     string
	Profession       string
	Description      string
	Value            float64
	Deadline         int16
	DenyContent      string
	UserID           string
	SocialNetworksID string
	BankAccountID    string
	CreatedAt        time.Time
}

//NewIdol create a new idol
func NewIdol(
	artisticName,
	profession,
	description string,
	value float64,
	deadline int16,
	denyContent,
	userID,
	socialNetworksID string,
	bankAccountID string,
) (*Idol, error) {
	i := &Idol{
		ID:               NewID(),
		ArtisticName:     strings.TrimSpace(artisticName),
		Profession:       strings.TrimSpace(profession),
		Description:      description,
		Value:            value,
		Deadline:         deadline,
		DenyContent:      denyContent,
		UserID:           userID,
		SocialNetworksID: socialNetworksID,
		BankAccountID:    bankAccountID,
		CreatedAt:        time.Now(),
	}

	err := i.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return i, nil
}

//Validate validate data
func (u *Idol) Validate() error {
	rgx, _ := regexp.Compile("^[A-Za-z]*$")

	if u.ArtisticName == "" {
		return ErrInvalidEntity
	}

	if u.Profession == "" || !rgx.MatchString(u.Profession) {
		return ErrInvalidEntity
	}

	if u.Description == "" {
		return ErrInvalidEntity
	}

	if u.Value == 0 {
		return ErrInvalidEntity
	}

	return nil
}
