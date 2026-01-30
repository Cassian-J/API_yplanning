package models

import (
	"errors"
	"net/http"
	"time"
)

type AvailabilityRequest struct {
	DateBegin time.Time `json:"date_begin"`
	DateEnd   time.Time `json:"date_end"`
	UserID    uint      `json:"user_id"`
}

func (a *AvailabilityRequest) Bind(r *http.Request) error {
	if a.DateBegin.IsZero() {
		return errors.New("date_begin must not be null")
	} else if a.DateEnd.IsZero() {
		return errors.New("date_end must not be null")
	} else if a.UserID < 1 {
		return errors.New("user_id must be >= 1")
	}
	return nil
}

type AvailabilityResponse struct {
	ID        uint      `json:"id"`
	DateBegin time.Time `json:"date_begin"`
	DateEnd   time.Time `json:"date_end"`
	UserID    uint      `json:"user_id"`
}
