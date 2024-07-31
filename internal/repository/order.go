package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/nadiannis/evento-api-fr/internal/domain"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) IOrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) GetAll() ([]*domain.Order, error) {
	query := "SELECT id, customer_id, ticket_id, quantity, total_price, created_at FROM orders"

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*domain.Order, 0)
	for rows.Next() {
		var order domain.Order

		err = rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.TicketID,
			&order.Quantity,
			&order.TotalPrice,
			&order.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) Add(order *domain.Order) error {
	query := `
	INSERT INTO orders (customer_id, ticket_id, quantity, total_price, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
`
	args := []any{order.CustomerID, order.TicketID, order.Quantity, order.TotalPrice, order.CreatedAt}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...).Scan(&order.ID)
}

func (r *OrderRepository) GetByCustomerID(customerID int64) ([]*domain.Order, error) {
	query := `
		SELECT id, customer_id, ticket_id, quantity, total_price, created_at
		FROM orders
		WHERE customer_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, customerID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make([]*domain.Order, 0)
	for rows.Next() {
		var order domain.Order

		err := rows.Scan(
			&order.ID,
			&order.CustomerID,
			&order.TicketID,
			&order.Quantity,
			&order.TotalPrice,
			&order.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		orders = append(orders, &order)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) DeleteByID(orderID int64) error {
	return nil
}

func (r *OrderRepository) DeleteAll() {
}
