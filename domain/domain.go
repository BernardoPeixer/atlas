package domain

import "context"

type CardsUseCase interface {
	// ListAllCards function to list all available cards
	ListAllCards(
		ctx context.Context,
	) error
}
