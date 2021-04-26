package user

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.User, error)
}

//Writer user writer
type Writer interface {
	Create(e *entity.User) (*entity.User, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	Get(id entity.ID) (*entity.User, error)
	Create(
		firstName string,
		lastName string,
		email string,
		phone string,
		genre string,
		category string,
	) (*entity.User, error)
}
