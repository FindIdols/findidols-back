package socialnetworks

import (
	"fmt"

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

//CreateSocialNetworks create a socialnetworks
func (s *Service) CreateSocialNetworks(
	youtube string,
	instagram string,
	twitter string,
	tiktok string,
) (entity.ID, error) {

	socialNetworks, err := entity.NewSocialNetworks(youtube, instagram, twitter, tiktok)

	if err != nil {
		fmt.Println("Erro ao criar redes sociais")
		return socialNetworks.ID, err
	}

	return s.repo.Create(socialNetworks)
}

//GetSocialNetworks get a idol
func (s *Service) GetSocialNetworks(id entity.ID) (*entity.SocialNetworks, error) {
	socialNetworks, err := s.repo.GetSocialNetworks(id)

	if socialNetworks == nil {
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return socialNetworks, nil
}
