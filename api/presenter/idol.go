package presenter

import (
	"github.com/FindIdols/findidols-back/entity"
)

//Idol data
type Idol struct {
	ID           entity.ID            `json:"id"`
	ArtisticName string               `json:"artistic_name"`
	Description  string               `json:"description"`
	Prices       *entity.PriceContent `json:"prices"`
}

//IdolInformation data
type IdolInformation struct {
	ID           entity.ID            `json:"id"`
	ArtisticName string               `json:"artistic_name"`
	Description  string               `json:"description"`
	Deadline     int16                `json:"deadline"`
	DenyContent  string               `json:"denyContent"`
	Profession   string               `json:"profession"`
	Youtube      string               `json:"youtube"`
	Instagram    string               `json:"instagram"`
	Twitter      string               `json:"twitter"`
	Tiktok       string               `json:"tiktok"`
	Prices       *entity.PriceContent `json:"prices"`
	VideoURL     []*entity.Video      `json:"videos"`
}
