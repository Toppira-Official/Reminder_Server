package middlewares

import "go.uber.org/fx"

var Module = fx.Module(
	"middlewares",
	fx.Provide(
		fx.Annotate(
			ErrorHandler,
			fx.ResultTags(`name:"error_handler"`),
		),
	),
)
