package main

import (
	"github.com/Toppira-Official/backend/internal/domain/entities"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "internal/domain/repositories",
		Mode: gen.WithDefaultQuery |
			gen.WithQueryInterface,
	})

	g.ApplyBasic(
		entities.User{},
		entities.Reminder{},
	)

	g.Execute()
}
