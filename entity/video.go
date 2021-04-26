package entity

import "time"

//Video data
type Video struct {
	ID           ID
	URL          string
	IdolID       string
	Introduction bool
	CreatedAt    time.Time
}

//NewVideo create a new user
func NewVideo(idolID, url string) (*Video, error) {
	v := &Video{
		ID:        NewID(),
		IdolID:    idolID,
		CreatedAt: time.Now(),
	}

	err := v.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return v, nil
}

//Validate validate data
func (i *Video) Validate() error {
	// if i.Location == "" {
	// 	return ErrInvalidEntity
	// }

	return nil
}
