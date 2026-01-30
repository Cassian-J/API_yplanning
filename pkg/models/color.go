package models

import (
	"errors"
	"net/http"
)

type ColorRequest struct {
	HexCode string `json:"hex_code"`
	Name    string `json:"name"`
}

func (c *ColorRequest) Bind(r *http.Request) error {
	if c.HexCode == "" {
		return errors.New("hex_code must not be null")
	} else if c.Name == "" {
		return errors.New("name must not be null")
	}
	return nil
}

type ColorResponse struct {
	ID      uint   `json:"id"`
	HexCode string `json:"hex_code"`
	Name    string `json:"name"`
}
