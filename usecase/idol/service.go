package idol

import (
	"fmt"
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

//CreateIdol create a idol
func (s *Service) CreateIdol(
	artisticName string,
	profession string,
	description string,
	value float64,
	deadline int16,
	userID string,
	socialNetworksID string,
	bankAccountID string,
) (entity.ID, error) {
	idol, err := entity.NewIdol(artisticName, profession, description, value, deadline, "", userID, socialNetworksID, bankAccountID)

	if err != nil {
		fmt.Println("idolo invalido")
		return idol.ID, err
	}

	idol.Value = convertMicroMoney(value)

	return s.repo.Create(idol)
}

//GetIdol get a idol
func (s *Service) GetIdol(id entity.ID) (*entity.Idol, error) {
	idol, err := s.repo.GetIdol(id)

	idol.Value = convertDefaultMoney(idol.Value)

	if idol == nil {
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return idol, nil
}

//GetIdols get a idol
func (s *Service) GetIdols() ([]*entity.Idol, error) {
	idols, err := s.repo.GetIdols()

	for _, idol := range idols {
		idol.Value = convertDefaultMoney(idol.Value)
	}

	if err != nil {
		return nil, err
	}

	if len(idols) == 0 {
		return nil, entity.ErrNotFound
	}

	return idols, nil
}

func convertMicroMoney(value float64) float64 {
	return value * 1000000
}

func convertDefaultMoney(value float64) float64 {
	return math.Round(value) / 1000000
}
