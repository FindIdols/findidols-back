package order

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Order, error)
}

//Writer order writer
type Writer interface {
	Create(e *entity.Order) (*entity.Order, error)
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetOrder(id entity.ID) (*entity.Order, error)
	CreateOrder(
		user *entity.User,
		usage string,
		subject string,
		instruction string,
		termsOfUse bool,
		idolID string,
	) (*entity.Order, error)
}
