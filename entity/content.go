package entity

import "strings"

//Content data
type Content struct {
	Usage       string
	Subject     string
	Instruction string
}

//NewContent create a new content
func NewContent(usage, subject, instruction string) (*Content, error) {
	c := &Content{
		Usage:       strings.TrimSpace(usage),
		Subject:     strings.TrimSpace(subject),
		Instruction: instruction,
	}

	err := c.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return c, nil
}

//Validate validate data
func (u *Content) Validate() error {

	if u.Usage == "" {
		return ErrInvalidEntity
	}

	if u.Subject == "" {
		return ErrInvalidEntity
	}

	if u.Instruction == "" {
		return ErrInvalidEntity
	}

	return nil
}
