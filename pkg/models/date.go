package models

import (
	"errors"
	"net/http"
	"time"
)

type DateRequest struct {
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	DateBegin    time.Time `json:"date_begin"`
	DateEnd      time.Time `json:"date_end"`
	UserID       uint      `json:"user_id"`
	Private      bool      `json:"private"`
	RecurrenceID uint      `json:"recurrence_id"`
	ColorID      uint      `json:"color_id"`
}

type RangeRequest struct {
	DateBegin time.Time `json:"date_begin"`
	DateEnd   time.Time `json:"date_end"`
	UserID    uint      `json:"user_id"`
}

func (u *DateRequest) Bind(r *http.Request) error {
	if u.Title == "" {
		return errors.New("title must not be null")
	} else if u.DateBegin.IsZero() {
		return errors.New("date_begin must not be null")
	} else if u.DateEnd.IsZero() {
		return errors.New("date_end must not be null")
	} else if u.UserID < 1 {
		return errors.New("user_id must be >= 1")
	}
	return nil
}

type DateResponse struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	Body         string    `json:"body"`
	DateBegin    time.Time `json:"date_begin"`
	DateEnd      time.Time `json:"date_end"`
	UserID       uint      `json:"user_id"`
	Private      bool      `json:"private"`
	RecurrenceID uint      `json:"recurrence_id"`
	ColorID      uint      `json:"color_id"`
}
