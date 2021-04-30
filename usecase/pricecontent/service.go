package pricecontent

import (
	"math"

	"github.com/FindIdols/findidols-back/entity"
)

//Service book usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//GetPricePerContent get a idol
func (s *Service) GetPricePerContent(id entity.ID) (*entity.PriceContent, error) {
	priceContent, err := s.repo.GetPricePerContent(id)

	priceContent.General = convertDefaultMoney(priceContent.General)

	if priceContent == nil {
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return priceContent, nil
}

//GetPricesPerContent get a idol
func (s *Service) GetPricesPerContent() ([]*entity.PriceContent, error) {

	pricesContent, err := s.repo.GetPricesPerContent()

	if err != nil {
		return nil, err
	}

	for _, price := range pricesContent {
		price.General = convertDefaultMoney(price.General)
	}

	if len(pricesContent) == 0 {
		return nil, entity.ErrNotFound
	}

	return pricesContent, nil
}

func convertDefaultMoney(value float64) float64 {
	return math.Round(value) / 1000000
}

func convertMicroMoney(value float64) float64 {
	return value * 1000000
}
