package models

import (
	"errors"
	"net/http"
	"yplanning/database/dbmodel"
)

type DateRequest struct {
	Title        string `json:"title"`
	Body         string `json:"body"`
	DateBegin    int64  `json:"date_begin"`
	DateEnd      int64  `json:"date_end"`
	UserID       uint   `json:"user_id"`
	Private      bool   `json:"private"`
	RecurrenceID *uint  `json:"recurrence_id"`
	ColorID      *uint   `json:"color_id"`
}

func (u *DateRequest) Bind(r *http.Request) error {
	if u.Title == "" {
		return errors.New("title must not be null")
	} else if u.DateBegin == 0 {
		return errors.New("date_begin must not be null")
	} else if u.DateEnd == 0 {
		return errors.New("date_end must not be null")
	} else if u.UserID < 1 {
		return errors.New("user_id must be >= 1")
	}
	return nil
}

type DateResponse struct {
	ID           uint           `json:"id"`
	Title        string         `json:"title"`
	Body         string         `json:"body"`
	DateBegin    int64            `json:"date_begin"`
	DateEnd      int64            `json:"date_end"`
	User         *UserResponse  `json:"user"`
	Private      bool           `json:"private"`
	RecurrenceID *uint           `json:"recurrence_id"`
	Color        *ColorResponse `json:"color"`
}

func ToDateResponse(date *dbmodel.Date) *DateResponse {
	if date == nil {
		return nil
	}

	return &DateResponse{
		ID:           date.ID,
		Title:        date.Title,
		Body:         date.Body,
		DateBegin:    date.BeginTime,
		DateEnd:      date.EndTime,
		User:         ToUserResponse(date.User),
		Private:      date.Private,
		RecurrenceID: date.RecurrenceID,
		Color:        ToColorResponse(date.Color),
	}
}
