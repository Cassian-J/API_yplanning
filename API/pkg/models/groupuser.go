package models

import (
	"errors"
	"net/http"
)

type GroupUserRequest struct {
	UserID  uint `json:"user_id"`
	GroupID uint `json:"group_id"`
	ColorID uint `json:"color_id"`
}

func (g *GroupUserRequest) Bind(r *http.Request) error {
	if g.UserID == 0 {
		return errors.New("user_id is required")
	}
	if g.GroupID == 0 {
		return errors.New("group_id is required")
	}
	return nil
}

type GroupUserResponse struct {
	UserID  uint `json:"user_id"`
	GroupID uint `json:"group_id"`
	ColorID uint `json:"color_id"`
}