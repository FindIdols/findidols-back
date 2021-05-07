package order

import (
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

//CreateOrder create a order
func (s *Service) CreateOrder(
	userId string,
	usage string,
	subject string,
	instruction string,
	termsOfUse bool,
	idolID string,
) (*entity.Order, error) {

	c, err := entity.NewContent(usage, subject, instruction)
	o, err := entity.NewOrder(userId, *c, termsOfUse, idolID)

	if err != nil {
		return o, err
	}

	return s.repo.Create(o)
}

//GetOrder get a order
func (s *Service) GetOrder(id entity.ID) (*entity.Order, error) {
	o, err := s.repo.Get(id)

	if o == nil {
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return o, nil
}
