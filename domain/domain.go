package domain

import (
	"atlas/domain/entities"
	"context"
)

type CardsUseCase interface {
	// ListAllCards function to list all available cards
	ListAllCards(
		ctx context.Context,
	) ([]entities.FestivalCards, error)
}
