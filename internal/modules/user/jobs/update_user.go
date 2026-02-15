package jobs

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/user/usecase/input"
	apperrors "github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/queues"

	"github.com/hibiken/asynq"
)

const TypeUpdateUser = "user:update"

type UpdateUserJob interface {
	Process(ctx context.Context, t *asynq.Task) error
}

type updateUserJob struct {
	uc usecase.UpdateUserUsecase
}

func NewUpdateUserJob(uc usecase.UpdateUserUsecase) UpdateUserJob {
	return &updateUserJob{uc: uc}
}

func (j *updateUserJob) Process(ctx context.Context, t *asynq.Task) error {
	var p input.UpdateUserInput
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return asynq.SkipRetry
	}

	_, err := j.uc.Execute(ctx, &p)
	if err == nil {
		return nil
	}

	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		switch appErr.Code {
		case apperrors.ErrUserNotFound, apperrors.ErrUserInvalidData, apperrors.ErrUserAlreadyExists:
			return asynq.SkipRetry
		}
	}

	return err
}

func Register(mux *asynq.ServeMux, updateUser UpdateUserJob) {
	mux.HandleFunc(TypeUpdateUser, updateUser.Process)
}

func EnqueueUpdateUser(q *queues.Client, in *input.UpdateUserInput, opts ...asynq.Option) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}

	task := asynq.NewTask(TypeUpdateUser, b)

	_, err = q.Enqueue(task, opts...)
	return err
}
