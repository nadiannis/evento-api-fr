package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/utils"
)

type TicketTypeRepository struct {
	db *sql.DB
}

func NewTicketTypeRepository(db *sql.DB) ITicketTypeRepository {
	return &TicketTypeRepository{
		db: db,
	}
}

func (r *TicketTypeRepository) GetAll() ([]*domain.TicketType, error) {
	query := "SELECT id, name, price FROM ticket_types"

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

	ticketTypes := make([]*domain.TicketType, 0)
	for rows.Next() {
		var ticketType domain.TicketType

		err = rows.Scan(&ticketType.ID, &ticketType.Name, &ticketType.Price)
		if err != nil {
			return nil, err
		}

		ticketTypes = append(ticketTypes, &ticketType)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ticketTypes, nil
}

func (r *TicketTypeRepository) Add(ticketType *domain.TicketType) error {
	query := `
		INSERT INTO ticket_types (name, price)
		VALUES ($1, $2)
		RETURNING id
	`
	args := []any{ticketType.Name, ticketType.Price}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, args...).Scan(&ticketType.ID)
	if err != nil {
		switch {
		case err.Error() == `ERROR: duplicate key value violates unique constraint "ticket_types_name_key" (SQLSTATE 23505)`:
			return utils.ErrTicketTypeAlreadyExists
		default:
			return err
		}
	}

	return nil
}

func (r *TicketTypeRepository) GetByName(ticketTypeName domain.TicketTypeName) (*domain.TicketType, error) {
	query := `
		SELECT id, name, price
		FROM ticket_types
		WHERE name = $1
	`

	var ticketType domain.TicketType

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, ticketTypeName).Scan(
		&ticketType.ID,
		&ticketType.Name,
		&ticketType.Price,
	)
	if err != nil {
		return nil, err
	}

	return &ticketType, nil
}
