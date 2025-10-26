package usecases

import (
	"atlas/domain"
	"atlas/domain/entities"
	"atlas/infrastructure/datastore"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type userUseCase struct {
	repository datastore.UserRepository
	cfg        entities.Config
}

func NewUserUseCase(
	repository datastore.UserRepository,
	cfg entities.Config,
) domain.UserUseCase {
	return userUseCase{
		repository: repository,
		cfg:        cfg,
	}
}

func (u userUseCase) RegisterUser(
	ctx context.Context,
	user entities.UserInfo,
) error {
	// TODO: Actually, the user can not defines username
	suffix := uuid.New()
	shortSuffix := strings.Split(suffix.String(), "-")[0]
	user.Username = fmt.Sprintf("user_%s", shortSuffix)

	return u.repository.RegisterUser(ctx, user)
}

func (u userUseCase) CheckUser(
	ctx context.Context,
	walletAddress string,
) (bool, error) {
	if strings.TrimSpace(walletAddress) == "" {
		return false, errors.New("invalid wallet address")
	}

	return u.repository.CheckUser(ctx, walletAddress)
}
