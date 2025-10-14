package repository

import (
	"atlas/domain/entities"
	"atlas/infrastructure/datastore"
	"context"
	"database/sql"
)

type cardRepository struct {
	conn func() *sql.DB
	cfg  entities.Config
}

func NewCardsRepository(
	settings datastore.RepositorySettings,
	cfg entities.Config,
) datastore.CardsRepository {
	return cardRepository{
		conn: settings.Connection,
		cfg:  cfg,
	}
}

func (c cardRepository) ListAllCards(
	ctx context.Context,
) error {
	return nil
}
