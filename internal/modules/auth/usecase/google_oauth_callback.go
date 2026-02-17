package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"

	"golang.org/x/oauth2"
)

type GoogleUserInfo struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	FamilyName string `json:"family_name"`
	Picture    string `json:"picture"`
}

type GoogleOauthCallbackUsecase interface {
	Execute(ctx context.Context, code, state, expectedState string) (*GoogleUserInfo, error)
}

type googleOauthCallbackUsecase struct {
	googleOauthConfig *oauth2.Config
}

func NewGoogleOauthCallbackUsecase(googleOauthConfig *oauth2.Config) GoogleOauthCallbackUsecase {
	return &googleOauthCallbackUsecase{googleOauthConfig: googleOauthConfig}
}

func (uc *googleOauthCallbackUsecase) Execute(
	ctx context.Context,
	code, state, expectedState string,
) (*GoogleUserInfo, error) {
	if code == "" {
		return nil, apperrors.E(apperrors.ErrAuthInvalidToken)
	}
	if expectedState == "" || state != expectedState {
		return nil, apperrors.E(apperrors.ErrAuthInvalidToken)
	}

	tok, err := uc.googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, apperrors.E(apperrors.ErrAuthInvalidToken, err)
	}

	client := uc.googleOauthConfig.Client(ctx, tok)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://openidconnect.googleapis.com/v1/userinfo", nil)
	if err != nil {
		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, apperrors.E(apperrors.ErrAuthInvalidToken)
	}

	var out GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, apperrors.E(apperrors.ErrServerNotResponding, err)
	}

	if out.Email == "" {
		return nil, apperrors.E(apperrors.ErrAuthInvalidToken)
	}

	return &out, nil
}
