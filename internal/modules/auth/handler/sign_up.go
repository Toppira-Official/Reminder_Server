package handler

import (
	"net/http"

	userUsecase "github.com/Toppira-Official/backend/internal/modules/user/usecase"
	"github.com/gin-gonic/gin"
)

type SignUpHandler struct {
	createUserUsecase userUsecase.CreateUserUsecase
}

func NewSignUpHandler(createUserUsecase userUsecase.CreateUserUsecase) *SignUpHandler {
	return &SignUpHandler{createUserUsecase: createUserUsecase}
}

func (hl *SignUpHandler) SignUpWithEmailPassword(c *gin.Context) {
	var input SignUpWithEmailPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, SignUpWithEmailPasswordOutput{Message: "invalid body request"})
		return
	}

	user, err := hl.createUserUsecase.Execute(c.Request.Context(), input.MapUser())
	if err != nil {
		c.JSON(http.StatusInternalServerError, SignUpWithEmailPasswordOutput{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, SignUpWithEmailPasswordOutput{Message: "welcome",
		Data: map[string]any{
			"user": user,
		}})
}
