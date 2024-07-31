package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nadiannis/evento-api-fr/internal/domain"
	"github.com/nadiannis/evento-api-fr/internal/utils"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) IEventRepository {
	return &EventRepository{
		db: db,
	}
}

func (r *EventRepository) GetAll() ([]*domain.Event, error) {
	query := "SELECT id, name, date FROM events"

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

	events := make([]*domain.Event, 0)
	for rows.Next() {
		var event domain.Event

		err := rows.Scan(&event.ID, &event.Name, &event.Date)
		if err != nil {
			return nil, err
		}

		events = append(events, &event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *EventRepository) Add(event *domain.Event) error {
	query := `
		INSERT INTO events (name, date)
		VALUES ($1, $2)
		RETURNING id
	`
	args := []any{event.Name, event.Date}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	return stmt.QueryRowContext(ctx, args...).Scan(&event.ID)
}

func (r *EventRepository) GetByID(eventID int64) (*domain.Event, error) {
	query := `
		SELECT id, name, date
		FROM events
		WHERE id = $1
	`

	var event domain.Event

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx, eventID).Scan(
		&event.ID,
		&event.Name,
		&event.Date,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, utils.ErrEventNotFound
		default:
			return nil, err
		}
	}

	return &event, nil
}

func (r *EventRepository) AddTicket(eventID int64, ticket *domain.Ticket) (*domain.Ticket, error) {
	return nil, nil
}
