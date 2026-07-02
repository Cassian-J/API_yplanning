package models

import (
	"errors"

	"net/http"

	"yplanning/database/dbmodel"
)

type GroupUserRequest struct {
	Username string `json:"username"`
	GroupID  uint   `json:"group_id"`
	ColorID  uint   `json:"color_id"`
}

func (g *GroupUserRequest) Bind(r *http.Request) error {
	if g.Username == "" {
		return errors.New("username is required")
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

	return &GroupUserResponse{
		User:  ToUserResponse(&userGroup.User),
		Group: ToGroupResponse(&userGroup.Group),
		Color: ToColorResponse(&userGroup.Color),
	}
}
