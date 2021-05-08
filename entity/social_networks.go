package entity

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

//SocialNetworks data
type SocialNetworks struct {
	ID        ID
	Youtube   string
	Twitch    string
	Nimo      string
	Instagram string
	Twitter   string
	TikTok    string
	CreatedAt time.Time
}

//NewSocialNetworks create a new SocialNetworks
func NewSocialNetworks(youtube, instagram, twitter, tiktok string) (*SocialNetworks, error) {
	sn := &SocialNetworks{
		ID:        NewID(),
		Youtube:   strings.TrimSpace(youtube),
		Instagram: strings.TrimSpace(instagram),
		Twitter:   strings.TrimSpace(twitter),
		TikTok:    strings.TrimSpace(tiktok),
		CreatedAt: time.Now(),
	}

	err := sn.Validate()

	if err != nil {
		fmt.Println("erro validacao")
		return nil, ErrInvalidEntity
	}

	return sn, nil
}

//Validate validate data
func (u *SocialNetworks) Validate() error {
	rgx, _ := regexp.Compile("youtube.com")

	if u.Youtube != "" && !rgx.MatchString(u.Youtube) {
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("instagram.com")

	if u.Instagram != "" && !rgx.MatchString(u.Instagram) {
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("twitter.com")

	if u.Twitter != "" && !rgx.MatchString(u.Twitter) {
		return ErrInvalidEntity
	}

	rgx, _ = regexp.Compile("tiktok.com")

	if u.TikTok != "" && !rgx.MatchString(u.TikTok) {
		return ErrInvalidEntity
	}

	return nil
}
