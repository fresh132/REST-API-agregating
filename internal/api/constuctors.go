package api

import "github.com/fresh132/REST-API-agregating/internal/repository"

type Handler struct {
	repo *repository.SubscriptionRepository
}

func NewHandler(repo *repository.SubscriptionRepository) *Handler {
	return &Handler{repo: repo}
}
