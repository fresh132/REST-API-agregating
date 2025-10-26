package validation

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type SubscriptionValidation struct {
	ServiceName string `json:"service_name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
}

func Validate(sv *SubscriptionValidation) error {
	if sv.Price <= 0 {
		return errors.New("price must be greater than 0")
	}

	if _, err := uuid.Parse(sv.UserID); err != nil {
		return errors.New("invalid user ID format")
	}

	startDate, err := time.Parse("01-2006", sv.StartDate)
	if err != nil {
		return errors.New("invalid start date format, use MM-YYYY")
	}

	if sv.EndDate == "" {
		return nil
	}

	endDate, err := time.Parse("01-2006", sv.EndDate)
	if err != nil {
		return errors.New("invalid end date format, use MM-YYYY")
	}

	if endDate.Before(startDate) {
		return errors.New("end_date cannot be before start_date")
	}

	return nil
}
