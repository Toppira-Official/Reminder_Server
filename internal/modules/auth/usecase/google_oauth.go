package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"github.com/Toppira-Official/backend/internal/configs"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOauthUsecase interface {
	Execute(ctx context.Context) (redirectUrl, state string)
}

type googleOauthUsecase struct {
	envs configs.Environments
}

func NewGoogleOauthUsecase(envs configs.Environments) GoogleOauthUsecase {
	return &googleOauthUsecase{envs: envs}
}

func (uc *googleOauthUsecase) Execute(ctx context.Context) (redirectUrl, state string) {
	googleOauthConfig := &oauth2.Config{
		RedirectURL:  uc.envs.GOOGLE_REDIRECT_URL.String(),
		ClientID:     uc.envs.GOOGLE_CLIENT_ID.String(),
		ClientSecret: uc.envs.GOOGLE_CLIENT_SECRET.String(),
		Scopes: []string{
			"email",
			"openid",
			"profile",
		},
		Endpoint: google.Endpoint,
	}

	state = generateState()
	redirectUrl = googleOauthConfig.AuthCodeURL(state)

	return
}

func generateState() string {
	b := make([]byte, 16)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}
