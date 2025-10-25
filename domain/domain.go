package domain

import (
	"atlas/domain/entities"
	"context"
)

type CardsUseCase interface {
	// ListAllCards function to list all available cards
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

// UserUseCase defines the interface for user-related business logic.
type UserUseCase interface {
	// RegisterUser handles the business logic to register a new user.
	// Receives the request context and user information.
	// Returns an error if the registration fails.
	RegisterUser(
		ctx context.Context,
		user entities.UserInfo,
	) error
}
