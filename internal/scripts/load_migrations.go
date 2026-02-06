package scripts

import (
	"github.com/Toppira-Official/backend/internal/domain/entities"
	"gorm.io/gorm"
)

func LoadMigrations(db *gorm.DB) error {
	return db.AutoMigrate(
		&entities.User{},
		&entities.Reminder{},
	)
}
