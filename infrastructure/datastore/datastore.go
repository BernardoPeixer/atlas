package datastore

import (
	"atlas/domain/entities"
	"atlas/util"
	"context"
	"database/sql"
)

// RepositorySettings defines the interface for managing database connections and server time.
type RepositorySettings interface {
	// Connection returns a database connection using the provided company configuration.
	// Returns a pointer to sql.DB.
	Connection() *sql.DB

	// Dismount closes all active connections to the database.
	// Returns an error if closing the connections fails.
	Dismount() error

	// ServerTime retrieves the current time from the internal server.
	// Accepts a context for cancellation and timeouts.
	// Returns a pointer to util.DateTime and an error if the operation fails.
	ServerTime(ctx context.Context) (*util.DateTime, error)
}

// CardsRepository defines the interface for managing festival cards and cryptocurrency types.
type CardsRepository interface {
	// ListAllCards retrieves all available festival cards.
	// Returns a slice of FestivalCards and an error if the operation fails.
	ListAllCards(
		ctx context.Context,
	) ([]entities.FestivalCard, error)

	// ListAllCryptoType retrieves all available cryptocurrency types.
	// Returns a slice of CryptoType and an error if the operation fails.
	ListAllCryptoType(
		ctx context.Context,
	) ([]entities.CryptoType, error)

	// RegisterCard saves a new festival card in the repository.
	// Accepts a FestivalCards entity and returns an error if the operation fails.
	RegisterCard(
		ctx context.Context,
		festivalCard entities.FestivalCard,
	) error
}

// UserRepository defines the interface for user-related persistence operations.
type UserRepository interface {
	// RegisterUser saves a new user in the repository.
	// Receives the request context and user information.
	// Returns an error if the registration fails.
	RegisterUser(
		ctx context.Context,
		user entities.UserInfo,
	) error

	// CheckUser verifies whether a user exists based on the provided wallet address.
	// Returns true if the user exists, or false otherwise, along with an error if the check fails.
	CheckUser(
		ctx context.Context,
		walletAddress string,
	) (bool, error)
}
