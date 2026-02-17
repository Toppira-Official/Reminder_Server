package handler

import (
	"net/http"

	"github.com/Toppira-Official/Reminder_Server/internal/configs"
	authUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"
	output "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"
	_ "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"

	"github.com/gin-gonic/gin"
)

type GoogleOauthHandler struct {
	googleOauthRedirectURLUsecase authUsecase.GoogleOauthRedirectURLUsecase
	googleOauthCallbackUsecase    authUsecase.GoogleOauthCallbackUsecase
	envs                          configs.Environments
}

func NewGoogleOauthHandler(
	googleOauthRedirectURLUsecase authUsecase.GoogleOauthRedirectURLUsecase,
	googleOauthCallbackUsecase authUsecase.GoogleOauthCallbackUsecase,
	envs configs.Environments,
) *GoogleOauthHandler {
	return &GoogleOauthHandler{
		googleOauthRedirectURLUsecase: googleOauthRedirectURLUsecase,
		googleOauthCallbackUsecase:    googleOauthCallbackUsecase,
		envs:                          envs,
	}
}

// GetGoogleOauthRedirectURL godoc
//
//	@Summary	Redirect to Google OAuth URL
//	@Tags		Authentication
//	@Success	307	{string}	string	Redirect	to	Google	OAuth	URL
//	@Failure	500	{object}	errors.ClientError
//	@Router		/auth/google-oauth/redirect-url [get]
func (h *GoogleOauthHandler) GetGoogleOauthRedirectURL(c *gin.Context) {
	ctx := c.Request.Context()
	redirectUrl, err := h.googleOauthRedirectURLUsecase.Execute(ctx)
	if err != nil {
		c.Error(err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, redirectUrl)
}

// GoogleOauthCallback godoc
//
//	@Summary	Handle Google OAuth callback
//	@Tags		Authentication
//	@Param		code	query		string	true	"Code"
//	@Param		state	query		string	true	"State"
//	@Success	200		{object}	output.HttpOutput
//	@Failure	401		{object}	errors.ClientError
//	@Failure	500		{object}	errors.ClientError
//	@Router		/auth/google-oauth/callback [get]
func (h *GoogleOauthHandler) GoogleOauthCallback(c *gin.Context) {
	ctx := c.Request.Context()

	code := c.Query("code")
	state := c.Query("state")

	userInfo, err := h.googleOauthCallbackUsecase.Execute(ctx, code, state)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, output.HttpOutput{
		Data: map[string]any{
			"user": userInfo,
		},
	})
}
