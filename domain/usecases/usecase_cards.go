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
) ([]entities.FestivalCard, error) {
	return c.repository.ListAllCards(ctx)
}

func (c cardsUseCase) ListAllCryptoType(
	ctx context.Context,
) ([]entities.CryptoType, error) {
	return c.repository.ListAllCryptoType(ctx)
}

func (c cardsUseCase) RegisterCard(
	ctx context.Context,
	festivalCard entities.FestivalCard,
) error {
	// TODO: Validate user wallet to register card
	return c.repository.RegisterCard(ctx, festivalCard)
}
