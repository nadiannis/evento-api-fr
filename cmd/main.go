package main

import (
	"database/sql"
	"flag"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/nadiannis/evento-api-fr/internal/handler"
	"github.com/nadiannis/evento-api-fr/internal/repository"
	"github.com/nadiannis/evento-api-fr/internal/usecase"
	"github.com/rs/zerolog/log"
)

type config struct {
	port int
	db   struct {
		driver string
		dsn    string
	}
}

type application struct {
	config   config
	handlers handler.Handlers
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.StringVar(&cfg.db.driver, "db-driver", "pgx", "PostgreSQL driver")
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:pass1234@localhost:5432/evento", "PostgreSQL data source name")

	flag.Parse()

	db, err := openDB(cfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer db.Close()
	log.Info().Msg("connected to database successfully")

	repos := repository.NewRepositories(db)
	usecases := usecase.NewUsecases(repos)
	handlers := handler.NewHandlers(usecases)

	app := &application{
		config:   cfg,
		handlers: handlers,
	}

	log.Info().Msg("add ticket types")
	prepopulateTicketTypes(usecases.TicketTypes)

	log.Info().Msg("add events and tickets")
	prepopulateEventsAndTickets(usecases.Events, usecases.Tickets)

	err = app.serve()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open(cfg.db.driver, cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
