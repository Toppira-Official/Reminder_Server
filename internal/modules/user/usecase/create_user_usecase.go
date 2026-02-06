package usecase

import (
	"context"
	"errors"

	"github.com/Toppira-Official/backend/internal/domain/constants"
	"github.com/Toppira-Official/backend/internal/domain/entities"
	"github.com/Toppira-Official/backend/internal/domain/repositories"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

type CreateUserUsecase interface {
	Execute(ctx context.Context, input *entities.User) (*entities.User, error)
}

type createUserUsecase struct {
	repo   *repositories.Query
	logger *zap.Logger
}

func NewCreateUserUsecase(repo *repositories.Query, logger *zap.Logger) CreateUserUsecase {
	return &createUserUsecase{repo: repo, logger: logger}
}

func (uc *createUserUsecase) Execute(ctx context.Context, input *entities.User) (*entities.User, error) {
	user := input
	err := uc.repo.User.WithContext(ctx).Save(user)
	if err != nil {
		uc.logger.Error("failed to create user", zap.Error(err))

		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, ErrUserAlreadyExists
		}

		return nil, constants.ErrInternalServer
	}

	return user, nil
}
