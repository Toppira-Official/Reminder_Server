package scripts

import (
	"go.uber.org/fx"
)

var Module = fx.Module(
	"scripts",
	fx.Invoke(
		LoadMigrations,
	),
)
