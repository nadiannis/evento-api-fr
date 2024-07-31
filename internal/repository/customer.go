package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/utils"
)

type CustomerRepository struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) ICustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) GetAll() ([]*domain.Customer, error) {
	query := "SELECT id, username, balance FROM customers"

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

	customers := make([]*domain.Customer, 0)
	for rows.Next() {
		var customer domain.Customer

		err = rows.Scan(&customer.ID, &customer.Username, &customer.Balance)
		if err != nil {
			return nil, err
		}

		customers = append(customers, &customer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (r *CustomerRepository) Add(customer *domain.Customer) error {
	query := `
		INSERT INTO customers (username, balance)
		VALUES ($1, $2)
		RETURNING id
	`
	args := []any{customer.Username, customer.Balance}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, args...).Scan(&customer.ID)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "customers_username_key" (SQLSTATE 23505)`:
			return utils.ErrCustomerAlreadyExists
		default:
			return err
		}
	}

	return nil
}

func (r *CustomerRepository) GetByID(customerID int64) (*domain.Customer, error) {
	query := `
		SELECT id, username, balance 
		FROM customers 
		WHERE id = $1
	`

	var customer domain.Customer

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, customerID).Scan(
		&customer.ID,
		&customer.Username,
		&customer.Balance,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrCustomerNotFound
		default:
			return nil, err
		}
	}

	return &customer, nil
}

func (r *CustomerRepository) AddOrder(customerID int64, order *domain.Order) error {
	return nil
}

func (r *CustomerRepository) DeleteAllOrders() {
}

func (r *CustomerRepository) AddBalance(customerID int64, amount float64) (*domain.Customer, error) {
	query := `
		UPDATE customers
		SET balance = balance + $1
		WHERE id = $2
		RETURNING id, username, balance
	`
	args := []any{amount, customerID}

	var customer domain.Customer

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, args...).Scan(
		&customer.ID,
		&customer.Username,
		&customer.Balance,
	)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (r *CustomerRepository) DeductBalance(customerID int64, amount float64) error {
	query := `
		UPDATE customers
		SET balance = balance - $1
		WHERE id = $2
		RETURNING id, username, balance
	`
	args := []any{amount, customerID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		switch {
		case err.Error() == `ERROR: new row for relation "customers" violates check constraint "customers_balance_check" (SQLSTATE 23514)`:
			return utils.ErrInsufficientBalance
		default:
			return err
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return utils.ErrCustomerNotFound
	}

	return nil
}
