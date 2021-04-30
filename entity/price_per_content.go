package entity

import (
	"fmt"
	"time"
)

//SocialNetworks data
type PriceContent struct {
	ID               ID
	General          float64
	Publicity        float64
	Politics         float64
	ProfessionalTips float64
	Earnings         float64
	CreatedAt        time.Time
}

//NewSocialNetworks create a new SocialNetworks
func NewPriceContent(generalPrice, publicityPrice, politicsPrice, professionalTipsPrice, earningsPrice float64) (*PriceContent, error) {
	pc := &PriceContent{
		ID:               NewID(),
		General:          generalPrice,
		Publicity:        publicityPrice,
		Politics:         politicsPrice,
		ProfessionalTips: professionalTipsPrice,
		Earnings:         earningsPrice,
		CreatedAt:        time.Now(),
	}

	err := pc.Validate()

	if err != nil {
		fmt.Println("erro validacao preco conteudo")
		return nil, ErrInvalidEntity
	}

	return pc, nil
}

//Validate validate data
func (pc *PriceContent) Validate() error {

	if pc.General <= 20 {
		return ErrInvalidEntity
	}

	return nil
}
