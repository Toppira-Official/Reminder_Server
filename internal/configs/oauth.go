package configs

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetGoogleOauthConfig(envs Environments) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  envs.GOOGLE_REDIRECT_URL.String(),
		ClientID:     envs.GOOGLE_CLIENT_ID.String(),
		ClientSecret: envs.GOOGLE_CLIENT_SECRET.String(),
		Scopes: []string{
			"openid",
			"email",
			"profile",
		},
		Endpoint: google.Endpoint,
	}
}
