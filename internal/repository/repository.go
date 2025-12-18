package repository

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Subscription struct {
	ID          int        `json:"id"`
	ServiceName string     `json:"service_name"`
	Price       int        `json:"price"`
	UserID      uuid.UUID  `json:"user_id"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type SubscriptionRepository struct {
	db *pgxpool.Pool
}

func NewSubscriptionRepository(db *pgxpool.Pool) *SubscriptionRepository {
	return &SubscriptionRepository{db: db}
}

func (r *SubscriptionRepository) Create(ctx context.Context, sub Subscription) (int, error) {
	var id int
	err := r.db.QueryRow(ctx,
		`INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
		 VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate,
	).Scan(&id)
	return id, err
}

func (r *SubscriptionRepository) GetByID(ctx context.Context, id int) (Subscription, error) {
	var sub Subscription

	err := r.db.QueryRow(ctx,
		`SELECT id, service_name, price, user_id, start_date, end_date, created_at, updated_at
		 FROM subscriptions WHERE id = $1`, id,
	).Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID,
		&sub.StartDate, &sub.EndDate, &sub.CreatedAt, &sub.UpdatedAt)
	return sub, err
}

func (r *SubscriptionRepository) Update(ctx context.Context, id int, sub Subscription) error {
	tag, err := r.db.Exec(
		ctx, `UPDATE subscriptions SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5, updated_at=NOW() WHERE id=$6`,
		sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate, id,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("subscription not found")
	}

	return nil
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id int) error {

	var exists bool
	err := r.db.QueryRow(ctx, `SELECT EXISTS(SELECT 1 FROM subscriptions WHERE id=$1)`, id).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("id not found")
	}

	_, err = r.db.Exec(ctx, `DELETE FROM subscriptions WHERE id=$1`, id)
	return err
}

func (r *SubscriptionRepository) ListSubscriptionsFil(ctx context.Context, userID *uuid.UUID, serviceName *string, limit, offset int) ([]Subscription, error) {
	var subs []Subscription
	args := []interface{}{}
	argIndex := 1
	query := `SELECT id, service_name, price, user_id, start_date, end_date, created_at, updated_at 
	          FROM subscriptions WHERE 1=1`

	if userID != nil {
		query += ` AND user_id=$` + strconv.Itoa(argIndex)
		args = append(args, *userID)
		argIndex++
	}

	if serviceName != nil {
		query += ` AND service_name=$` + strconv.Itoa(argIndex)
		args = append(args, *serviceName)
		argIndex++
	}

	query += ` ORDER BY created_at DESC LIMIT $` + strconv.Itoa(argIndex) + ` OFFSET $` + strconv.Itoa(argIndex+1)
	args = append(args, limit, offset)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s Subscription
		if err := rows.Scan(
			&s.ID, &s.ServiceName, &s.Price, &s.UserID,
			&s.StartDate, &s.EndDate, &s.CreatedAt, &s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}

	return subs, nil
}

func (r *SubscriptionRepository) GetTotal(ctx context.Context, userID *uuid.UUID, serviceName *string, startDate, endDate time.Time) (int64, error) {
	var total int64
	args := []interface{}{}
	argIndex := 1

	// Calculate the number of months for each subscription within the period
	// Formula: (year_diff * 12) + month_diff + 1 (to include both start and end month)
	query := `
		SELECT COALESCE(SUM(price * (
			(EXTRACT(YEAR FROM LEAST(COALESCE(end_date, $1), $1)) - EXTRACT(YEAR FROM GREATEST(start_date, $2))) * 12 +
			(EXTRACT(MONTH FROM LEAST(COALESCE(end_date, $1), $1)) - EXTRACT(MONTH FROM GREATEST(start_date, $2))) + 1
		)), 0) AS total
		FROM subscriptions
		WHERE start_date <= $1
		AND (end_date IS NULL OR end_date >= $2)
	`
	args = append(args, endDate, startDate)
	argIndex = 3

	if userID != nil {
		query += ` AND user_id=$` + strconv.Itoa(argIndex)
		args = append(args, *userID)
		argIndex++
	}

	if serviceName != nil {
		query += ` AND service_name=$` + strconv.Itoa(argIndex)
		args = append(args, *serviceName)
		argIndex++
	}

	err := r.db.QueryRow(ctx, query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
