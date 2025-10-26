package api

import (
	"context"
	"net/http"
	"time"

	"github.com/fresh132/REST-API-agregating/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CalculateTotalCost godoc
// @Summary Calculate total cost
// @Description Calculate total cost of subscriptions for a specific period with optional filtering
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param user_id query string false "User ID for filtering"
// @Param service_name query string false "Service name for filtering"
// @Param start_date query string true "Start date in MM-YYYY format"
// @Param end_date query string true "End date in MM-YYYY format"
// @Success 200 {object} TotalCostResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/cost [get]
func (h *Handler) CalculateTotalCost(c *gin.Context) {
	var userID *uuid.UUID

	if userIDStr := c.Query("user_id"); userIDStr != "" {
		parsedUserID, err := uuid.Parse(userIDStr)
		if err != nil {
			logger.Warn.Warn("invalid user ID format",
				"user_id", userIDStr,
				"error", err.Error(),
			)
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID format"})
			return
		}

		userID = &parsedUserID
	}

	var serviceName *string

	if serviceNameStr := c.Query("service_name"); serviceNameStr != "" {
		serviceName = &serviceNameStr
	}

	startDateStr := c.Query("start_date")

	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		logger.Warn.Warn("start_date and end_date are required", "start_date", startDateStr, "end_date", endDateStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	startDate, err := time.Parse("01-2006", startDateStr)

	if err != nil {
		logger.Warn.Warn("invalid start_date format", "start_date", startDateStr, "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format, use MM-YYYY"})
		return
	}

	endDate, err := time.Parse("01-2006", endDateStr)

	if err != nil {
		logger.Warn.Warn("invalid end_date format", "end_date", endDateStr, "error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format, use MM-YYYY"})
		return
	}

	ctx := context.Background()

	total, err := h.repo.GetTotal(ctx, userID, serviceName, startDate, endDate)

	if err != nil {
		logger.Error.Error("failed to calculate total cost",
			"error", err.Error(),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to calculate total cost"})
		return
	}

	logger.Info.Info("Calculate Total Cost OK",
		"user_id", userID,
		"service_name", serviceName,
		"start_date", startDateStr,
		"end_date", endDateStr,
		"total_cost", total,
	)
	c.JSON(http.StatusOK, gin.H{"total_cost": total})
}
