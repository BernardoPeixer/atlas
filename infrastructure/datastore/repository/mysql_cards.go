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
) ([]entities.FestivalCard, error) {
	festivalCards := make([]entities.FestivalCard, 0)

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
	       u.status_code,
	       u.created_at,
	       u.modified_at,
	       c.id,
	       c.name,
	       c.symbol,
	       c.status_code,
	       c.created_at,
	       c.modified_at
	FROM festival_cards fc
	    INNER JOIN cryptos c ON c.id = fc.id_crypto_type 
	    INNER JOIN user u ON u.id = fc.id_user 
	ORDER BY fc.status_code, fc.created_at DESC 
	`

	rows, err := c.conn().QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error in queryContext (query): %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var festivalCard entities.FestivalCard
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
			&festivalCard.UserInfo.StatusCode,
			&festivalCard.UserInfo.CreatedAt,
			&festivalCard.UserInfo.ModifiedAt,
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

func (c cardRepository) ListAllCryptoType(
	ctx context.Context,
) ([]entities.CryptoType, error) {
	cryptoTypes := make([]entities.CryptoType, 0)

	// language=sql
	query := `
	SELECT c.id,
	       c.name,
	       c.symbol,
	       c.status_code,
	       c.created_at,
	       c.modified_at
	FROM cryptos c
	WHERE status_code = 0
	`

	rows, err := c.conn().QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error in queryContext (query): %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cryptoType entities.CryptoType
		err = rows.Scan(
			&cryptoType.ID,
			&cryptoType.Name,
			&cryptoType.Symbol,
			&cryptoType.StatusCode,
			&cryptoType.CreatedAt,
			&cryptoType.ModifiedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error in scan (rows): %v", err)
		}

		cryptoTypes = append(cryptoTypes, cryptoType)
	}

	return cryptoTypes, nil
}

func (c cardRepository) RegisterCard(
	ctx context.Context,
	festivalCard entities.FestivalCard,
) error {
	// language=sql
	query := `
	INSERT INTO festival_cards(id_crypto_type, id_user, balance, crypto_price) 
	VALUES (?, ?, ?, ?)
	`

	_, err := c.conn().ExecContext(
		ctx,
		query,
		festivalCard.CryptoType.ID,
		festivalCard.UserInfo.ID,
		festivalCard.Balance,
		festivalCard.CryptoPrice,
	)
	if err != nil {
		return fmt.Errorf("error in execContext (query): %v", err)
	}

	return nil
}

func (c cardRepository) FinishTransactionCard(
	ctx context.Context,
	cardID int64,
) error {
	// language=sql
	query := `
	UPDATE festival_cards 
	SET status_code = 2,
	    sold_at = CURRENT_TIMESTAMP
	WHERE id = ?
	`

	_, err := c.conn().ExecContext(ctx, query, cardID)
	if err != nil {
		return fmt.Errorf("error in execContext (query): %v", err)
	}

	return nil
}
