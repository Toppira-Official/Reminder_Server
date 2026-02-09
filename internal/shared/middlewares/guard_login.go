package middlewares

import (
	"strings"

	authUsecase "github.com/Toppira-Official/backend/internal/modules/auth/usecase"
	apperrors "github.com/Toppira-Official/backend/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

func GuardLogin(
	verifyJwtUsecase authUsecase.VerifyJwtUsecase,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			appErr := &apperrors.AppError{ClientError: apperrors.ClientError{Code: apperrors.ErrAuthTokenNotProvided}}
			c.JSON(apperrors.HTTPStatus(appErr.Code), appErr.Client())
			c.Abort()
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			appErr := &apperrors.AppError{ClientError: apperrors.ClientError{Code: apperrors.ErrAuthInvalidToken}}
			c.JSON(apperrors.HTTPStatus(appErr.Code), appErr.Client())
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := verifyJwtUsecase.Execute(c, token)
		if err != nil {
			appErr := &apperrors.AppError{ClientError: apperrors.ClientError{Code: apperrors.ErrAuthInvalidToken}}
			c.JSON(apperrors.HTTPStatus(appErr.Code), appErr.Client())
			c.Abort()
			return
		}

		c.Set("userClaims", claims)

		c.Next()
	}
}
