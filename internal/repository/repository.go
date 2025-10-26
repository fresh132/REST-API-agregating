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
	_, err := r.db.Exec(
		ctx, `UPDATE subscriptions SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5, updated_at=NOW() WHERE id=$6`,
		sub.ServiceName, sub.Price, sub.UserID, sub.StartDate, sub.EndDate, id,
	)

	return err
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

func (r *SubscriptionRepository) ListSubscriptionsFil(ctx context.Context, userID *uuid.UUID, serviceName *string) ([]Subscription, error) {
	query := `SELECT id, service_name, price, user_id, start_date, end_date, created_at, updated_at 
	          FROM subscriptions WHERE 1=1`

	args := []interface{}{}
	argIndex := 1

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

	query += ` ORDER BY created_at DESC`

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subs []Subscription
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

func (r *SubscriptionRepository) GetTotal(ctx context.Context, userID *uuid.UUID, serviceName *string, startDate, endDate time.Time) (int, error) {
	query := `
		SELECT COALESCE(SUM(price), 0)
		FROM subscriptions
		WHERE (start_date <= $2)
		AND (end_date IS NULL OR end_date >= $1)
	`
	args := []interface{}{startDate, endDate}
	argIndex := 3

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

	var total int
	err := r.db.QueryRow(ctx, query, args...).Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}
