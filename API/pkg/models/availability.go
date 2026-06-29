package models

import (
	"errors"
	"net/http"
	"yplanning/database/dbmodel"
)

type AvailabilityRequest struct {
	DateBegin int  `json:"date_begin"`
	DateEnd   int  `json:"date_end"`
	UserID    uint `json:"user_id"`
}

func (a *AvailabilityRequest) Bind(r *http.Request) error {
	if a.DateBegin == 0 {
		return errors.New("date_begin must not be null")
	} else if a.DateEnd == 0 {
		return errors.New("date_end must not be null")
	} else if a.UserID < 1 {
		return errors.New("user_id must be >= 1")
	}
	return nil
}

type AvailabilityResponse struct {
	ID        uint          `json:"id"`
	DateBegin int           `json:"date_begin"`
	DateEnd   int           `json:"date_end"`
	User      *UserResponse `json:"user"`
}

func ToAvailabilityResponse(availability *dbmodel.Availability) *AvailabilityResponse {

	if availability == nil {
		return nil
	}

	return &AvailabilityResponse{
		ID:        availability.ID,
		DateBegin: availability.BeginTime,
		DateEnd:   availability.EndTime,
		User:      ToUserResponse(availability.User),
	}
}
