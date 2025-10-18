package repository

import (
	"atlas/domain/entities"
	"atlas/infrastructure/datastore"
	"context"
	"database/sql"
	"fmt"
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
) ([]entities.FestivalCards, error) {
	festivalCards := make([]entities.FestivalCards, 0)

	// language=sql
	query := `
	SELECT fc.id,
	       fc.balance,
	       fc.crypto_price,
	       fc.status_code,
	       fc.created_at,
	       fc.modified_at_,
	       fc.sold_at,
	       u.id,
	       u.username,
	       u.wallet_address,
	       c.id,
	       c.name,
	       c.symbol,
	       c.status_code,
	       c.created_at,
	       c.modified_at
	FROM festival_cards fc
	    INNER JOIN cryptos c ON c.id = fc.id_crypto_type 
	    INNER JOIN user u ON u.id = fc.id_user
	WHERE fc.status_code != 2
	`

	rows, err := c.conn().QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error in queryContext (query): %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var festivalCard entities.FestivalCards
		err = rows.Scan(
			&festivalCard.ID,
			&festivalCard.Balance,
			&festivalCard.CryptoPrice,
			&festivalCard.StatusCode,
			&festivalCard.CreatedAt,
			&festivalCard.ModifiedAt,
			&festivalCard.SoldAt,
			&festivalCard.UserInfo.ID,
			&festivalCard.UserInfo.Username,
			&festivalCard.UserInfo.WalletAddress,
			&festivalCard.CryptoType.ID,
			&festivalCard.CryptoType.Name,
			&festivalCard.CryptoType.Symbol,
			&festivalCard.CryptoType.StatusCode,
			&festivalCard.CryptoType.CreatedAt,
			&festivalCard.CryptoType.ModifiedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error in scan (rows): %v", err)
		}

		festivalCards = append(festivalCards, festivalCard)
	}

	return festivalCards, nil
}
