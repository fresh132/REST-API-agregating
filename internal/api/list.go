package api

import (
	"context"
	"net/http"

	"github.com/fresh132/REST-API-agregating/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ListSubscriptions godoc
// @Summary List subscriptions
// @Description Get a list of subscriptions with optional filtering
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param user_id query string false "User ID for filtering"
// @Param service_name query string false "Service name for filtering"
// @Success 200 {array} repository.Subscription
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions [get]
func (h *Handler) ListSubscriptions(c *gin.Context) {
	var userID *uuid.UUID

	if userIDStr := c.Query("user_id"); userIDStr != "" {
		parsedUserID, err := uuid.Parse(userIDStr)

		if err != nil {
			logger.Error.Error("invalid user ID format",
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

	ctx := context.Background()

	subscriptions, err := h.repo.ListSubscriptionsFil(ctx, userID, serviceName)
	if err != nil {
		logger.Error.Error("failed to get subscriptions",
			"error", err.Error(),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get subscriptions",
			"details": err.Error(),
		})
		return
	}

	logger.Info.Info("Get list OK", "id", userID)
	c.JSON(http.StatusOK, subscriptions)
}
