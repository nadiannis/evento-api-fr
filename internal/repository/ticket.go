package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/utils"
)

type TicketRepository struct {
	db *sql.DB
}

func NewTicketRepository(db *sql.DB) ITicketRepository {
	return &TicketRepository{
		db: db,
	}
}

func (r *TicketRepository) GetAll() ([]*domain.Ticket, error) {
	query := "SELECT id, event_id, type, quantity FROM tickets"

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

	tickets := make([]*domain.Ticket, 0)
	for rows.Next() {
		var ticket domain.Ticket

		err := rows.Scan(&ticket.ID, &ticket.EventID, &ticket.Type, &ticket.Quantity)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, &ticket)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *TicketRepository) Add(ticket *domain.Ticket) error {
	query := `
		SELECT EXISTS(
			SELECT 1
			FROM tickets
			WHERE event_id = $1 AND type = $2
		)
	`
	args := []any{ticket.EventID, ticket.Type}

	var exists bool

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	checkStmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer checkStmt.Close()

	err = checkStmt.QueryRowContext(ctx, args...).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		return utils.ErrTicketAlreadyExists
	}

	query = `
		INSERT INTO tickets (event_id, type, quantity)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	args = []any{ticket.EventID, ticket.Type, ticket.Quantity}

	insertStmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRowContext(ctx, args...).Scan(&ticket.ID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *TicketRepository) GetByID(ticketID int64) (*domain.Ticket, error) {
	query := `
		SELECT id, event_id, type, quantity
		FROM tickets
		WHERE id = $1
	`

	var ticket domain.Ticket

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, ticketID).Scan(
		&ticket.ID,
		&ticket.EventID,
		&ticket.Type,
		&ticket.Quantity,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrTicketNotFound
		default:
			return nil, err
		}
	}

	return &ticket, nil
}

func (r *TicketRepository) GetByEventID(eventID int64) ([]*domain.Ticket, error) {
	query := `
		SELECT id, event_id, type, quantity
		FROM tickets
		WHERE event_id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tickets := make([]*domain.Ticket, 0)
	for rows.Next() {
		var ticket domain.Ticket

		err := rows.Scan(&ticket.ID, &ticket.EventID, &ticket.Type, &ticket.Quantity)
		if err != nil {
			return nil, err
		}

		tickets = append(tickets, &ticket)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tickets, nil
}

func (r *TicketRepository) AddQuantity(ticketID int64, quantity int) (*domain.Ticket, error) {
	query := `
		UPDATE tickets
		SET quantity = quantity + $1
		WHERE id = $2
		RETURNING id, event_id, type, quantity
	`

	var ticket domain.Ticket

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, quantity, ticketID).Scan(
		&ticket.ID,
		&ticket.EventID,
		&ticket.Type,
		&ticket.Quantity,
	)

	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (r *TicketRepository) DeductQuantity(ticketID int64, quantity int) error {
	query := `
		UPDATE tickets
		SET quantity = quantity - $1
		WHERE id = $2
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, quantity, ticketID)
	if err != nil {
		switch {
		case err.Error() == `ERROR: new row for relation "tickets" violates check constraint "tickets_quantity_check" (SQLSTATE 23514)`:
			return utils.ErrInsufficientTicketQuantity
		default:
			return err
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return utils.ErrTicketNotFound
	}

	return nil
}
