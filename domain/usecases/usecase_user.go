package usecases

import (
	"atlas/domain"
	"atlas/domain/entities"
	"atlas/infrastructure/datastore"
	"context"
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

	return u.RegisterUser(ctx, user)
}
