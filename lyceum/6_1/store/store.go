package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	pool *pgxpool.Pool
}

func New(ctx context.Context, connString string) (*Store, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("create pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return &Store{pool: pool}, nil
}

func (s *Store) Close() {
	s.pool.Close()
}

func (s *Store) Pool() *pgxpool.Pool {
	return s.pool
}

func (s *Store) CreateOrder(ctx context.Context, item string, quantity int) (int, error) {
	var id int
	err := s.pool.QueryRow(ctx, "INSERT INTO orders (item, quantity) VALUES ($1, $2) RETURNING id", item, quantity).Scan(&id)
	return id, err
}

func (s *Store) GetOrder(ctx context.Context, id int) (string, int, error) {
	var item string
	var quantity int
	err := s.pool.QueryRow(ctx, "SELECT item, quantity FROM orders WHERE id = $1", id).Scan(&item, &quantity)
	return item, quantity, err
}

func (s *Store) UpdateOrder(ctx context.Context, id int, item string, quantity int) error {
	_, err := s.pool.Exec(ctx, "UPDATE orders SET item = $1, quantity = $2 WHERE id = $3", item, quantity, id)
	return err
}

func (s *Store) DeleteOrder(ctx context.Context, id int) error {
	_, err := s.pool.Exec(ctx, "DELETE FROM orders WHERE id = $1", id)
	return err
}

func (s *Store) ListOrders(ctx context.Context) ([]Order, error) {
	rows, err := s.pool.Query(ctx, "SELECT id, item, quantity FROM orders ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []Order
	for rows.Next() {
		var o Order
		if err := rows.Scan(&o.ID, &o.Item, &o.Quantity); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, rows.Err()
}

type Order struct {
	ID       int
	Item     string
	Quantity int
}
