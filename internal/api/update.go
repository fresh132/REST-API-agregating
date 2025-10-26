package api

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/fresh132/REST-API-agregating/internal/logger"
	"github.com/fresh132/REST-API-agregating/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UpdateSubscription godoc
// @Summary Update subscription
// @Description Update an existing subscription by ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Param subscription body UpdateSubscriptionRequest true "Updated subscription data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [put]
func (h *Handler) UpdateSubscription(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		logger.Error.Error("Invalid subscription ID",
			"id", idStr,
			"error", err.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subscription ID"})
		return
	}

	var input struct {
		ServiceName string `json:"service_name" binding:"required"`
		Price       int    `json:"price" binding:"required"`
		UserID      string `json:"user_id" binding:"required"`
		StartDate   string `json:"start_date" binding:"required"`
		EndDate     string `json:"end_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	startDate, err := time.Parse("01-2006", input.StartDate)

	if err != nil {
		logger.Warn.Warn("Invalid start date format",
			"start_date", input.StartDate,
			"error", err.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start date format"})
		return
	}

	var endDate *time.Time

	if input.EndDate != "" {
		parsedEndDate, err := time.Parse("01-2006", input.EndDate)
		if err != nil {
			logger.Error.Error("Invalid end date format",
				"end_date", input.EndDate,
				"error", err.Error(),
			)
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end date format"})
			return
		}

		endDate = &parsedEndDate
	}

	userUUID, err := uuid.Parse(input.UserID)

	if err != nil {
		logger.Error.Error("Invalid user ID format",
			"user_id", input.UserID,
			"error", err.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID format"})
		return
	}

	sub := repository.Subscription{
		ServiceName: input.ServiceName,
		Price:       input.Price,
		UserID:      userUUID,
		StartDate:   startDate,
		EndDate:     endDate,
		UpdatedAt:   time.Now(),
	}

	ctx := context.Background()

	err = h.repo.Update(ctx, id, sub)

	if err != nil {
		logger.Error.Error("Failed to update subscription",
			"id", id,
			"error", err.Error(),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update subscription"})
		return
	}

	logger.Info.Info("subscription updated successfully",
		"id", id,
		"service_name", input.ServiceName,
		"price", input.Price,
		"user_id", input.UserID,
		"start_date", input.StartDate,
		"end_date", input.EndDate,
	)
	c.JSON(http.StatusOK, gin.H{"message": "subscription updated successfully"})
}
