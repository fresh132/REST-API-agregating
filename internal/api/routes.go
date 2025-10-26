package api

import (
	"github.com/fresh132/REST-API-agregating/internal/repository"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, repo *repository.SubscriptionRepository) {
	handler := NewHandler(repo)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/subscriptions", handler.CreateSubscription)
	router.GET("/subscriptions/:id", handler.GetSubscriptions)
	router.PUT("/subscriptions/:id", handler.UpdateSubscription)
	router.DELETE("/subscriptions/:id", handler.DeleteSubscription)
	router.GET("/subscriptions", handler.ListSubscriptions)
	router.GET("/subscriptions/cost", handler.CalculateTotalCost)
}
