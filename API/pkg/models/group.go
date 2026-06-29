package models

import (
	"errors"
	"net/http"
	"yplanning/database/dbmodel"
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
	ID      uint         `json:"id"`
	Name    string       `json:"name"`
	Creator *UserResponse `json:"creator"`
}

func ToGroupResponse(group *dbmodel.Group) *GroupResponse {
	if group == nil {
		return nil
	}

	return &GroupResponse{
		ID:      group.ID,
		Name:    group.Name,
		Creator: ToUserResponse(group.Creator),
	}
}
