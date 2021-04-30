package pricecontent

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Reader interface
type Reader interface {
	GetPricePerContent(id entity.ID) (*entity.PriceContent, error)
	GetPricesPerContent() ([]*entity.PriceContent, error)
}

//Writer idol writer
type Writer interface {
}

//Repository interface
type Repository interface {
	Reader
	Writer
}

//UseCase interface
type UseCase interface {
	GetPricePerContent(id entity.ID) (*entity.PriceContent, error)
	GetPricesPerContent() ([]*entity.PriceContent, error)
}
