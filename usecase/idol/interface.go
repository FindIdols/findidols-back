package idol

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Reader interface
type Reader interface {
	GetIdol(id entity.ID) (*entity.Idol, error)
	GetIdols() ([]*entity.Idol, error)
}

//Writer idol writer
type Writer interface {
	Create(e *entity.Idol) (entity.ID, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	CreateIdol(
		artisticName string,
		profession string,
		description string,
		value float64,
		deadline int16,
		userID string,
		socialNetworksID string,
		bankAccountID string,
	) (entity.ID, error)
	GetIdol(idolID entity.ID) (*entity.Idol, error)
	GetIdols() ([]*entity.Idol, error)
}
