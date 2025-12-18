package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// SubscriptionRepo defines methods for subscription repository operations
type SubscriptionRepo interface {
	Create(ctx context.Context, sub Subscription) (int, error)
	GetByID(ctx context.Context, id int) (Subscription, error)
	Update(ctx context.Context, id int, sub Subscription) error
	Delete(ctx context.Context, id int) error
	ListSubscriptionsFil(ctx context.Context, userID *uuid.UUID, serviceName *string, limit, offset int) ([]Subscription, error)
	GetTotal(ctx context.Context, userID *uuid.UUID, serviceName *string, startDate, endDate time.Time) (int64, error)
}
