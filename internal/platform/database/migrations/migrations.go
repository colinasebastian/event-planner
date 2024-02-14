package migrations

import (
	"goproject.com/eventplanner-io/api/internal/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{}, &models.Registration{}, &models.User{})
}
