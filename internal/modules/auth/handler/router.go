package handler

import "github.com/gin-gonic/gin"

func RegisterRoutes(engine *gin.Engine, handler *SignUpHandler) {
	group := engine.Group("/auth")

	group.POST("/sign-up-with-user-password", handler.SignUpWithEmailPassword)
}
