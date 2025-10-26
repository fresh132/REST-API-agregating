package api

import (
	"net/http"

	"github.com/fresh132/REST-API-agregating/internal/validation"
	"github.com/gin-gonic/gin"
)

func CreateSubscription(c *gin.Context) {
	var input struct {
		ServiceName string `json:"service_name" binding:"required"`
		Price       int    `json:"price" binding:"required"`
		UserID      string `json:"user_id" binding:"required"`
		StartDate   string `json:"start_date" binding:"required"`
		EndDate     string `json:"end_date"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if input.Price <= 0 {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

}
