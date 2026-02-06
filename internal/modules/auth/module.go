package auth

import (
	"github.com/Toppira-Official/backend/internal/modules/auth/handler"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"auth",
	fx.Provide(handler.NewSignUpHandler),
	fx.Invoke(handler.RegisterRoutes),
)
