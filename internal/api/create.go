package api

import (
	"context"
	"net/http"
	"time"

	"github.com/fresh132/REST-API-agregating/internal/logger"
	"github.com/fresh132/REST-API-agregating/internal/repository"
	"github.com/fresh132/REST-API-agregating/internal/validation"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateSubscription godoc
// @Summary Create a new subscription
// @Description Create a new subscription with the provided data
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body CreateSubscriptionRequest true "Subscription data"
// @Success 201 {object} CreateSubscriptionResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [post]
func (h *Handler) CreateSubscription(c *gin.Context) {
	var input struct {
		ServiceName string `json:"service_name" binding:"required"`
		Price       int    `json:"price" binding:"required"`
		UserID      string `json:"user_id" binding:"required"`
		StartDate   string `json:"start_date" binding:"required"`
		EndDate     string `json:"end_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		logger.Error.Error("Failed to bind JSON", "error", err.Error(), "UserID", input.UserID)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if input.Price <= 0 {
		logger.Error.Error("The price must be greater than 0.", "price", input.Price)
		c.JSON(http.StatusBadRequest, gin.H{"error": "The price must be greater than 0."})
		return
	}

	err := validation.Validate(&validation.SubscriptionValidation{
		ServiceName: input.ServiceName,
		Price:       input.Price,
		UserID:      input.UserID,
		StartDate:   input.StartDate,
		EndDate:     input.EndDate,
	})

	if err != nil {
		logger.Error.Error("Validation error", "error", err.Error(), "UserID", input.UserID)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userUUID, err := uuid.Parse(input.UserID)
	if err != nil {
		logger.Error.Error("Invalid user ID format", "error", err.Error(), "UserID", input.UserID)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID format"})
		return
	}

	startDate, err := time.Parse("01-2006", input.StartDate)
	if err != nil {
		logger.Error.Error("Invalid start date format", "error", err.Error(), "StartDate", input.StartDate)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start date format"})
		return
	}

	var endDate *time.Time
	if input.EndDate != "" {
		parsedEndDate, err := time.Parse("01-2006", input.EndDate)
		if err != nil {
			logger.Error.Error("Invalid end date format", "error", err.Error(), "EndDate", input.EndDate)
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end date format"})
			return
		}
		endDate = &parsedEndDate
	}

	sub := repository.Subscription{
		ServiceName: input.ServiceName,
		Price:       input.Price,
		UserID:      userUUID,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   time.Now(),
	}

	ctx := context.Background()

	id, err := h.repo.Create(ctx, sub)
	if err != nil {
		logger.Error.Error("Failed to create subscription", "error", err.Error(), "UserID", input.UserID)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create subscription"})
		return
	}

	logger.Info.Info("Create Subscription OK", "id", id, "UserID", input.UserID)
	c.JSON(http.StatusCreated, gin.H{
		"id":      id,
		"message": "subscription created successfully",
	})

}
