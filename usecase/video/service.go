package video

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Service video usecase
type Service struct {
	repo Repository
}

//NewService create new service
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Get(idolID entity.ID) (*entity.Video, error) {
	v, err := s.repo.Get(idolID)

	if v == nil {
		return nil, entity.ErrNotFound
	}

	if err != nil {
		return nil, err
	}

	return v, nil
}

func (s *Service) GetVideos(idolID entity.ID) ([]*entity.Video, error) {
	videos, err := s.repo.GetVideos(idolID)

	if err != nil {
		return nil, err
	}

	if len(videos) == 0 {
		return nil, entity.ErrNotFound
	}

	return videos, nil
}
