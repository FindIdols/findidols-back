package presenter

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Order data
type Order struct {
	ID          entity.ID `json:"id"`
	OrderNumber int       `json:"orderNumber"`
}
