package handler

import (
	"net/http"

	"github.com/Toppira-Official/Reminder_Server/internal/configs"
	authUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"
	"github.com/gin-gonic/gin"
)

const (
	googleOauthStateCookieName = "google_oauth_state"
)

type GoogleOauthHandler struct {
	googleOauthRedirectURLUsecase authUsecase.GoogleOauthRedirectURLUsecase
	envs                          configs.Environments
}

func NewGoogleOauthHandler(
	googleOauthRedirectURLUsecase authUsecase.GoogleOauthRedirectURLUsecase,
	envs configs.Environments,
) *GoogleOauthHandler {
	return &GoogleOauthHandler{
		googleOauthRedirectURLUsecase: googleOauthRedirectURLUsecase,
		envs:                          envs,
	}
}

// GetGoogleOauthRedirectURL godoc
//
//	@Summary	Redirect to Google OAuth URL
//	@Tags		Authentication
//	@Success	307	{string}	string	Redirect	to	Google	OAuth	URL
//	@Router		/auth/google-oauth/redirect-url [get]
func (h *GoogleOauthHandler) GetGoogleOauthRedirectURL(c *gin.Context) {
	ctx := c.Request.Context()
	redirectUrl := h.googleOauthRedirectURLUsecase.Execute(ctx)

	c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}
