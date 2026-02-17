package usecase

import (
	"context"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/oauth2"
)

type GoogleOauthRedirectURLUsecase interface {
	Execute(ctx context.Context) (redirectUrl, state string)
}

type googleOauthRedirectURLUsecase struct {
	googleOauthConfig *oauth2.Config
}

func NewGoogleOauthRedirectURLUsecase(googleOauthConfig *oauth2.Config) GoogleOauthRedirectURLUsecase {
	return &googleOauthRedirectURLUsecase{googleOauthConfig: googleOauthConfig}
}

func (uc *googleOauthRedirectURLUsecase) Execute(ctx context.Context) (redirectUrl, state string) {
	state = generateState()
	redirectUrl = uc.googleOauthConfig.AuthCodeURL(state)

	return
}

func generateState() string {
	b := make([]byte, 16)
	rand.Read(b)

	return base64.URLEncoding.EncodeToString(b)
}
