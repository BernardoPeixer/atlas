package repository

import (
	"atlas/domain/entities"
	"atlas/infrastructure/datastore"
	"context"
	"database/sql"
	"fmt"
)

type userRepository struct {
	conn func() *sql.DB
	cfg  entities.Config
}

func NewUserRepository(
	settings datastore.RepositorySettings,
	cfg entities.Config,
) datastore.UserRepository {
	return userRepository{
		conn: settings.Connection,
		cfg:  cfg,
	}
}

func (u userRepository) RegisterUser(
	ctx context.Context,
	user entities.UserInfo,
) error {
	// language=sql
	query := `
	INSERT INTO user (username, wallet_address) 
	VALUES(?, ?)
	`

	_, err := u.conn().ExecContext(ctx, query, user.Username, user.WalletAddress)
	if err != nil {
		return fmt.Errorf("error execContext (query): %v", err)
	}

	return nil
}
