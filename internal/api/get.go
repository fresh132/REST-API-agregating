package api

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSubscription godoc
// @Summary Get subscription by ID
// @Description Retrieve a specific subscription by its ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "Subscription ID"
// @Success 200 {object} repository.Subscription
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /subscriptions/{id} [get]
func (h *Handler) GetSubscriptions(c *gin.Context) {
	IDString := c.Param("id")

	id, err := strconv.Atoi(IDString)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid subscription ID"})
		return
	}

	ctx := context.Background()
	subscription, err := h.repo.GetByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	c.JSON(http.StatusOK, subscription)

}
