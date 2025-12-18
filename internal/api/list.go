package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fresh132/REST-API-agregating/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ListSubscriptions godoc
// @Summary List subscriptions
// @Description Get a list of subscriptions with optional filtering and pagination
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param user_id query string false "User ID for filtering"
// @Param service_name query string false "Service name for filtering"
// @Param limit query int false "Number of results per page (default 10)"
// @Param offset query int false "Number of results to skip (default 0)"
// @Success 200 {object} ListSubscriptionsResponse
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

	limit := 10
	offset := 0

	if limitStr := c.Query("limit"); limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		} else {
			logger.Warn.Warn("invalid limit parameter",
				"limit", limitStr,
			)
		}
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil && o >= 0 {
			offset = o
		} else {
			logger.Warn.Warn("invalid offset parameter",
				"offset", offsetStr,
			)
		}
	}

	ctx := context.Background()

	subscriptions, err := h.repo.ListSubscriptionsFil(ctx, userID, serviceName, limit, offset)
	if err != nil {
		logger.Error.Error("failed to get subscriptions",
			"error", err.Error(),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get subscriptions",
			"details": err.Error(),
		})
		return
	}

	logger.Info.Info("Get list OK", "user_id", userID, "limit", limit, "offset", offset)
	c.JSON(http.StatusOK, gin.H{
		"data":   subscriptions,
		"limit":  limit,
		"offset": offset,
	})
}
