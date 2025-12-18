package api

import "github.com/fresh132/REST-API-agregating/internal/repository"

type Handler struct {
	repo repository.SubscriptionRepo
}

func NewHandler(repo repository.SubscriptionRepo) *Handler {
	return &Handler{repo: repo}
}
