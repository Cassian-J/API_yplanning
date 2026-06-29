package models

import (
	"errors"
	"net/http"
	"yplanning/database/dbmodel"
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

func ToColorResponse(color *dbmodel.Color) *ColorResponse {
	if color == nil {
		return nil
	}

	return &ColorResponse{
		ID:      color.ID,
		Name:    color.Name,
		HexCode: color.HexCode,
	}
}
