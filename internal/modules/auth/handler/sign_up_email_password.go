package handler

import (
	"net/http"
	"strconv"

	authInput "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/handler/dto/input"
	authOutput "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/handler/dto/output"
	authUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/auth/usecase"
	userUsecase "github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase"
	userInput "github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase/input"
	sharedDto "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

type SignUpHandler struct {
	createUserUsecase   userUsecase.CreateUserUsecase
	hashPasswordUsecase authUsecase.HashPasswordUsecase
	generateJwtUsecase  authUsecase.GenerateJwtUsecase
}

func NewSignUpHandler(
	createUserUsecase userUsecase.CreateUserUsecase,
	hashPasswordUsecase authUsecase.HashPasswordUsecase,
	generateJwtUsecase authUsecase.GenerateJwtUsecase,
) *SignUpHandler {
	return &SignUpHandler{
		createUserUsecase:   createUserUsecase,
		hashPasswordUsecase: hashPasswordUsecase,
		generateJwtUsecase:  generateJwtUsecase,
	}
}

// SignUpWithEmailPassword godoc
//
//	@Summary	sign up with email and password
//	@Tags		Authentication
//	@Accept		json
//	@Produce	json
//	@Param		body	body		authInput.SignUpWithEmailPasswordInput	true	"Sign Up Input"
//	@Success	201		{object}	sharedDto.HttpOutput[authOutput.AuthOutput]
//	@Failure	400		{object}	apperrors.ClientError
//	@Failure	409		{object}	apperrors.ClientError
//	@Failure	500		{object}	apperrors.ClientError
//	@Failure	503		{object}	apperrors.ClientError
//	@Router		/auth/sign-up-with-user-password [post]
func (hl *SignUpHandler) SignUpWithEmailPassword(c *gin.Context) {
	ctx := c.Request.Context()

	var input authInput.SignUpWithEmailPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperrors.E(apperrors.ErrUserInvalidData, err))
		return
	}

	usecaseInput := &userInput.CreateUserInput{
		Email:    input.Email,
		Password: &input.Password,
		IsActive: false,
	}
	savedUser, err := hl.createUserUsecase.Execute(ctx, usecaseInput)
	if err != nil {
		c.Error(err)
		return
	}

	userIDString := strconv.Itoa(int(savedUser.ID))
	accessToken, err := hl.generateJwtUsecase.Execute(ctx, userIDString)
	if err != nil {
		c.Error(err)
		return
	}

	savedUser.Password = nil

	c.JSON(http.StatusCreated, sharedDto.HttpOutput[authOutput.AuthOutput]{
		Data: authOutput.AuthOutput{
			User:        sharedDto.ToUserOutput(savedUser),
			AccessToken: accessToken,
		},
	})
}
