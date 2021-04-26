package user

import (
	"fmt"

	"github.com/FindIdols/findidols-back/entity"
)

//Service user usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

//Create create a user
func (s *Service) Create(
	firstName string,
	lastName string,
	email string,
	phone string,
	genre string,
	category string,
) (*entity.User, error) {
	u, err := entity.NewUser(firstName, lastName, email, phone, genre, category)

	if err != nil {
		fmt.Println("erro ao montar entidade user")
		return u, err
	}

	return s.repo.Create(u)
}

//Get get a user
func (s *Service) Get(id entity.ID) (*entity.User, error) {
	u, err := s.repo.Get(id)

	if u == nil {
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return u, nil
}
