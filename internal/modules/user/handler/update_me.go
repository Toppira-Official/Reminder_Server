package handler

import (
	"errors"
	"net/http"

	dtoInput "github.com/Toppira-Official/Reminder_Server/internal/modules/user/handler/dto/input"
	userOutput "github.com/Toppira-Official/Reminder_Server/internal/modules/user/handler/dto/output"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/jobs"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase"
	userUcInput "github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase/input"
	sharedDto "github.com/Toppira-Official/Reminder_Server/internal/shared/dto"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/entities"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/queues"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/sony/gobreaker/v2"
	"go.uber.org/zap"
)

type UpdateMeHandler struct {
	updateUserUsecase usecase.UpdateUserUsecase
	q                 *queues.Client
	logger            *zap.Logger
}

func NewUpdateMeHandler(updateUserUsecase usecase.UpdateUserUsecase, q *queues.Client, logger *zap.Logger) *UpdateMeHandler {
	return &UpdateMeHandler{
		updateUserUsecase: updateUserUsecase,
		q:                 q,
		logger:            logger,
	}
}

// UpdateMyInfo godoc
//
//	@Summary	update my(user) data
//	@Tags		User
//	@Accept		json
//	@Produce	json
//	@Param		body	body		dtoInput.UpdateMeInput	true	"Update Me Input"
//	@Success	200		{object}	sharedDto.HttpOutput[userOutput.UpdateMeOutput]
//	@Success	202		{object}	sharedDto.HttpOutput[userOutput.UpdateMeAcceptedOutput]
//	@Failure	400		{object}	apperrors.ClientError
//	@Failure	401		{object}	apperrors.ClientError
//	@Failure	500		{object}	apperrors.ClientError
//	@Failure	503		{object}	apperrors.ClientError
//	@Security	BearerAuth
//	@Router		/user/me [patch]
func (hl *UpdateMeHandler) UpdateMyInfo(c *gin.Context) {
	ctx := c.Request.Context()

	userVal, exists := c.Get("user")
	if !exists {
		c.Error(apperrors.E(apperrors.ErrUserNotFound))
		return
	}

	user, ok := userVal.(*entities.User)
	if !ok {
		c.Error(apperrors.E(apperrors.ErrUserNotFound))
		return
	}

	var input dtoInput.UpdateMeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apperrors.E(apperrors.ErrUserInvalidData, err))
		return
	}

	usecaseInput := &userUcInput.UpdateUserInput{
		ID:       user.ID,
		Name:     input.Name,
		Password: input.Password,
		Phone:    input.Phone,
	}
	updatedUser, err := hl.updateUserUsecase.Execute(ctx, usecaseInput)
	if err == nil {
		updatedUser.Password = nil
		c.JSON(http.StatusOK, sharedDto.HttpOutput[userOutput.UpdateMeOutput]{
			Data: userOutput.UpdateMeOutput{
				User: sharedDto.ToUserOutput(updatedUser),
			},
		})

		return
	}

	if errors.Is(err, gobreaker.ErrOpenState) {
		if enqErr := jobs.EnqueueUpdateUser(
			hl.q,
			usecaseInput,
			asynq.Queue("critical"),
			asynq.MaxRetry(10),
			asynq.ProcessIn(usecase.UpdateUserRetryTime),
		); enqErr != nil {
			hl.logger.Error("failed to enqueue update user task", zap.Error(enqErr))
			c.Error(apperrors.E(apperrors.ErrServiceTemporarilyUnavailable, enqErr))
			return
		}
		c.JSON(http.StatusAccepted, sharedDto.HttpOutput[userOutput.UpdateMeAcceptedOutput]{
			Data: userOutput.UpdateMeAcceptedOutput{
				Message: "Update queued for processing",
			},
		})
		return
	}
}
