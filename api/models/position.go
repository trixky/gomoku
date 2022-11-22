package models

import "errors"

type Position struct {
	X uint8 `json:"x"`
	Y uint8 `json:"y"`
}

// Sanitize sanitizes these attributes
func (p *Position) Sanitize() error {
	if p.X >= 19 {
		return errors.New("the value of X must be between 0 and 18")
	}
	if p.Y >= 19 {
		return errors.New("the value of Y must be between 0 and 18")
	}
	return nil
}
