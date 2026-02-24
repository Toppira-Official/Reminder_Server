package notification

import (
	"github.com/Toppira-Official/Reminder_Server/internal/modules/notification/adapters"
	"github.com/Toppira-Official/Reminder_Server/internal/modules/notification/providers"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"notification",
	fx.Provide(
		providers.GetFirebase,
		fx.Annotate(
			adapters.NewFirebaseAdaptor,
			fx.ResultTags(`name:"firebase_adaptor"`),
		),
	),
)
