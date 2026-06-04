package models

import (
	"errors"
	"net/http"
)

type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	ColorID  uint   `json:"color_id"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	ColorID  *uint  `json:"color_id,omitempty"`
}

type GetUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (g *GetUserRequest) Bind(r *http.Request) error {
	panic("unimplemented")
}

func (u *UserRequest) Bind(r *http.Request) error {
	if u.Email == "" {
		return errors.New("email must not be null")
	} else if u.Password == "" {
		return errors.New("password must not be null")
	} else if u.Username == "" {
		return errors.New("username must not be null")
	}
	return nil
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	ColorID  *uint  `json:"color_id,omitempty"`
}
