package repositories

import (
	"goproject.com/eventplanner-io/api/internal/models"
	"gorm.io/gorm"
)

type EventsRepository interface {
	GetAll() ([]models.Event, error)
	GetByID(id int64) (*models.Event, error)
	Save(e *models.Event) error
	Update(e *models.Event) error
	Delete(e models.Event) error
}

type eventsrepository struct {
	db *gorm.DB
}

func NewEventsRepository(db *gorm.DB) *eventsrepository {
	return &eventsrepository{
		db: db,
	}
}

func (er *eventsrepository) GetAll() ([]models.Event, error) {
	var events []models.Event
	err := er.db.Preload("User").Find(&events).Error

	if err != nil {
		return nil, err
	}
	return events, nil
}

func (er *eventsrepository) GetByID(id int64) (*models.Event, error) {
	var event models.Event
	err := er.db.Preload("User").Model(&event).Where("id = ?", id).FirstOrInit(&event).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (er *eventsrepository) Save(event *models.Event) error {
	return er.db.Create(&event).Error
}

func (er *eventsrepository) Update(event *models.Event) error {
	return er.db.Model(&event).Updates(models.Event{
		Name:        event.Name,
		Description: event.Description,
		Location:    event.Location,
		DateTime:    event.DateTime,
		UserID:      event.UserID,
		User:        event.User,
	}).Error
}

func (er *eventsrepository) Delete(event models.Event) error {
	return er.db.Delete(&event).Error
}
