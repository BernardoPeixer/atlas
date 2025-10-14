package datastore

import (
	"atlas/util"
	"context"
	"database/sql"
)

type RepositorySettings interface {
	// Connection returns a database connection by the provided company config
	Connection() *sql.DB

	// Dismount closes all connections with database
	Dismount() error

	// ServerTime returns the time in internal server
	ServerTime(ctx context.Context) (*util.DateTime, error)
}

type CardsRepository interface {
	// ListAllCards function to list all available cards
	ListAllCards(
		ctx context.Context,
	) error
}
