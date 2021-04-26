package socialnetworks

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Reader interface
type Reader interface {
	GetSocialNetworks(id entity.ID) (*entity.SocialNetworks, error)
}

//Writer idol writer
type Writer interface {
	Create(e *entity.SocialNetworks) (entity.ID, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	CreateSocialNetworks(
		youtube string,
		instagram string,
		twitter string,
		tiktok string,
	) (entity.ID, error)
	GetSocialNetworks(id entity.ID) (*entity.SocialNetworks, error)
}
