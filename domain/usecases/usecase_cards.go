package usecases

import (
	"atlas/domain"
	"atlas/domain/entities"
	"atlas/infrastructure/datastore"
	"context"
)

type cardsUseCase struct {
	repository datastore.CardsRepository
	cfg        entities.Config
}

func NewCardsUseCase(
	repository datastore.CardsRepository,
	cfg entities.Config,
) domain.CardsUseCase {
	return cardsUseCase{
		repository: repository,
		cfg:        cfg,
	}
}

func (c cardsUseCase) ListAllCards(
	ctx context.Context,
) ([]entities.FestivalCards, error) {
	return c.repository.ListAllCards(ctx)
}
