package models

import (
	"errors"
	"net/http"
	"yplanning/database/dbmodel"
)

type AvailabilityRequest struct {
	DateBegin int64 `json:"dateBegin"`
	DateEnd   int64 `json:"dateEnd"`
	UserID    uint  `json:"userId"`
}

func (a *AvailabilityRequest) Bind(r *http.Request) error {
	if a.DateBegin == 0 {
		return errors.New("dateBegin must not be null")
	} else if a.DateEnd == 0 {
		return errors.New("dateEnd must not be null")
	} else if a.UserID < 1 {
		return errors.New("userId must be >= 1")
	}
	return nil
}

type AvailabilityResponse struct {
	ID        uint          `json:"id"`
	DateBegin int64         `json:"dateBegin"`
	DateEnd   int64         `json:"dateEnd"`
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
