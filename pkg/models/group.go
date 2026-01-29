package models

import (
	"errors"
	"net/http"
)

type GroupRequest struct {
	Name      string `json:"name" binding:"required"`
	CreatorID uint   `json:"creator_id" binding:"required"`
}

func (a *GroupRequest) Bind(r *http.Request) error {
	if a.Name == "" {
		return errors.New("name must not be null")
	} else if a.CreatorID == 0 {
		return errors.New("invalid creator ID")
	}
	return nil
}

type GroupResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatorID uint   `json:"creator_id"`
}
