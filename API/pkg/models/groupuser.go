package models

import (
	"errors"

	"net/http"

	"yplanning/database/dbmodel"
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
	User  *UserResponse  `json:"user"`
	Group *GroupResponse `json:"group"`
	Color *ColorResponse `json:"color"`
}

func ToGroupUserResponse(userGroup *dbmodel.UserGroup) *GroupUserResponse {
	if userGroup == nil {
		return nil
	}

	var color *ColorResponse
	if userGroup.Color.ID != 0 {
		color = ToColorResponse(&userGroup.Color)
	}

	return &GroupUserResponse{
		User:  ToUserResponse(&userGroup.User),
		Group: ToGroupResponse(&userGroup.Group),
		Color: color,
	}
}
