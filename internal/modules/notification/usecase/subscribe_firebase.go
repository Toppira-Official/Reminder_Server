package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Toppira-Official/Reminder_Server/internal/modules/notification/usecase/input"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"

	"github.com/Toppira-Official/Reminder_Server/internal/shared/repositories"
	"github.com/sony/gobreaker/v2"
	"gorm.io/gorm"
)

const SubscribeFirebaseRetryTime = 30 * time.Second

type SubscribeFirebaseUsecase interface {
	Execute(ctx context.Context, input *input.SubscribeFirebaseInput) (*entities.FirebaseSubscriber, error)
}

type subscribeFirebaseUsecase struct {
	repo    *repositories.Query
	breaker *gobreaker.CircuitBreaker[struct{}]
}

func NewSubscribeFirebaseUsecase(repo *repositories.Query) SubscribeFirebaseUsecase {
	settings := gobreaker.Settings{
		Name:        "subscribe_firebase_db",
		MaxRequests: 1,
		Interval:    0,
		Timeout:     SubscribeFirebaseRetryTime,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures >= 3
		},
		IsSuccessful: func(err error) bool {
			if err == nil {
				return true
			}
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return true
			}
			return false
		},
	}

	return &subscribeFirebaseUsecase{repo: repo, breaker: gobreaker.NewCircuitBreaker[struct{}](settings)}
}

func (s *subscribeFirebaseUsecase) Execute(ctx context.Context, input *input.SubscribeFirebaseInput) (*entities.FirebaseSubscriber, error) {
	firebaseSubscriber := &entities.FirebaseSubscriber{
		Token:  input.Token,
		UserID: input.UserID,
	}

	_, err := s.breaker.Execute(func() (struct{}, error) {
		return struct{}{}, s.repo.FirebaseSubscriber.WithContext(ctx).Create(firebaseSubscriber)
	})

	if err != nil {
		if errors.Is(err, gobreaker.ErrOpenState) {
			return nil, apperrors.E(apperrors.ErrServiceTemporarilyUnavailable, err)
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, apperrors.E(apperrors.ErrUserAlreadyExists, err)
		}
		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return firebaseSubscriber, nil
}
