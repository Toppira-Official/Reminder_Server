package adapters

import (
	"context"
	"time"

	"firebase.google.com/go/messaging"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/notification/domain/contract"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/notification/domain/model"
	"github.com/Toppira-Official/Reminder_Server/internal/shared/errors"
	"github.com/avast/retry-go"
)

type FirebaseAdaptor struct {
	messaging *messaging.Client
}

func NewFirebaseAdaptor(messaging *messaging.Client) contract.Sender {
	return &FirebaseAdaptor{messaging: messaging}
}

func (f *FirebaseAdaptor) Send(ctx context.Context, message model.Message) error {
	if message.Token == nil {
		return errors.E(errors.ErrServerInternalError)
	}

	notification := &messaging.Notification{
		Title: message.Title,
		Body:  message.Body,
	}
	if message.ImageURL != nil {
		notification.ImageURL = *message.ImageURL
	}

	err := retry.Do(
		func() error {
			_, err := f.messaging.Send(ctx, &messaging.Message{
				Token:        *message.Token,
				Notification: notification,
			})
			return err
		},
		retry.Attempts(5),
		retry.DelayType(retry.CombineDelay(retry.BackOffDelay, retry.RandomDelay)),
		retry.Delay(1*time.Second),
		retry.MaxJitter(500*time.Millisecond),
		retry.LastErrorOnly(true),
	)

	if err != nil {
		return errors.E(errors.ErrServerInternalError)
	}

	return nil
}
