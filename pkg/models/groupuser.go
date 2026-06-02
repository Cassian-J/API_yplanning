package models

import "net/http"

type GroupUserRequest struct {
	UserID  uint `json:"user_id"`
	GroupID uint `json:"group_id"`
	ColorID uint `json:"color_id"`
}

func (g *GroupUserRequest) Bind(r *http.Request) error {
	panic("unimplemented")
}

type GroupUserResponse struct {
	UserID  uint `json:"user_id"`
	GroupID uint `json:"group_id"`
	ColorID uint `json:"color_id"`
}
