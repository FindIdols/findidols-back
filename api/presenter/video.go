package presenter

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Video data
type Video struct {
	ID           entity.ID `json:"video_id"`
	VideoURL     string    `json:"videoUrl"`
	Introduction string    `json:"introduction"`
}
