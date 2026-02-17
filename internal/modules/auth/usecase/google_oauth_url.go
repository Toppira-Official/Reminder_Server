package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/redis/go-redis/v9"
	"golang.org/x/oauth2"
)

type GoogleOauthRedirectURLUsecase interface {
	Execute(ctx context.Context) (redirectUrl string, err error)
}

type googleOauthRedirectURLUsecase struct {
	googleOauthConfig *oauth2.Config
	cache             *redis.Client
}

func NewGoogleOauthRedirectURLUsecase(googleOauthConfig *oauth2.Config, cache *redis.Client) GoogleOauthRedirectURLUsecase {
	return &googleOauthRedirectURLUsecase{googleOauthConfig: googleOauthConfig, cache: cache}
}

func (uc *googleOauthRedirectURLUsecase) Execute(ctx context.Context) (string, error) {
	state := generateState()
	redirectUrl := uc.googleOauthConfig.AuthCodeURL(state)

	err := uc.cache.Set(ctx, fmt.Sprintf("oauth:state:%s", state), "1", 5*time.Minute).Err()
	if err != nil {
		return "", apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	return redirectUrl, nil
}

func generateState() string {
	b := make([]byte, 16)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}
