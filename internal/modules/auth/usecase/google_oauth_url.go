package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/oauth2"
)

type GoogleOauthRedirectURLUsecase interface {
	Execute(ctx context.Context) (redirectUrl string)
}

type googleOauthRedirectURLUsecase struct {
	googleOauthConfig *oauth2.Config
	cache             *redis.Client
}

func NewGoogleOauthRedirectURLUsecase(googleOauthConfig *oauth2.Config, cache *redis.Client) GoogleOauthRedirectURLUsecase {
	return &googleOauthRedirectURLUsecase{googleOauthConfig: googleOauthConfig, cache: cache}
}

func (uc *googleOauthRedirectURLUsecase) Execute(ctx context.Context) (redirectUrl string) {
	state := generateState()
	redirectUrl = uc.googleOauthConfig.AuthCodeURL(state)

	uc.cache.Set(ctx, fmt.Sprintf("oauth:state:%s", state), "1", 5*time.Minute)

	return redirectUrl
}

func generateState() string {
	b := make([]byte, 16)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}
