package presenter

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Idol data
type Idol struct {
	ID           entity.ID `json:"id"`
	ArtisticName string    `json:"artistic_name"`
	Description  string    `json:"description"`
	Value        float64   `json:"value"`
}

//IdolInformation data
type IdolInformation struct {
	ID           entity.ID       `json:"id"`
	ArtisticName string          `json:"artistic_name"`
	Description  string          `json:"description"`
	Value        float64         `json:"value"`
	Deadline     int16           `json:"deadline"`
	DenyContent  string          `json:"denyContent"`
	Profession   string          `json:"profession"`
	Youtube      string          `json:"youtube"`
	Instagram    string          `json:"instagram"`
	Twitter      string          `json:"twitter"`
	Tiktok       string          `json:"tiktok"`
	VideoURL     []*entity.Video `json:"videos"`
}
