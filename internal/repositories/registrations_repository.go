package repositories

import (
	"goproject.com/eventplanner-io/api/internal/models"
	"gorm.io/gorm"
)

type RegistrationsRepository interface {
	Save(r *models.Registration) error
	Delete(r models.Registration) error
}

type registrationsRepository struct {
	db *gorm.DB
}

func NewRegistrationsRepository(db *gorm.DB) *registrationsRepository {
	return &registrationsRepository{
		db: db,
	}
}

func (rr *registrationsRepository) Save(registration *models.Registration) error {
	return rr.db.Create(&registration).Error
}

func (rr *registrationsRepository) Delete(registration models.Registration) error {
	return rr.db.Where(&models.Registration{
		EventID: registration.EventID,
		UserID:  registration.UserID}).Delete(&registration).Error
}
