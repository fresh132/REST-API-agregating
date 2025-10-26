package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/fresh132/REST-API-agregating/internal/logger"
	"github.com/gin-gonic/gin"
)

// DeleteSubscription godoc
// @Summary Delete subscription
// @Description Delete a subscription by ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /subscriptions/{id} [delete]
func (h *Handler) DeleteSubscription(c *gin.Context) {
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

	ctx := context.Background()

	err = h.repo.Delete(ctx, id)

	if err != nil {
		logger.Error.Error("failed to delete subscription",
			"id", id,
			"error", err.Error(),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete subscription",
			"details": err.Error(),
		})
		return
	}

	logger.Info.Info("Delete Subscription OK", "id", idStr)
	c.JSON(http.StatusOK, gin.H{"message": "subscription deleted successfully"})
}
