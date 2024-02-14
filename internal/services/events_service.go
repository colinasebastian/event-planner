package services

import (
	"errors"

	"goproject.com/eventplanner-io/api/internal/models"
	"goproject.com/eventplanner-io/api/internal/repositories"
)

type EventsService interface {
	GetAll() ([]models.Event, error)
	GetByID(id int64) (*models.Event, error)
	Save(event *models.Event) error
	Update(event *models.Event) error
	Delete(event models.Event) error
}

type eventsService struct {
	repository repositories.EventsRepository
}

func NewEventsService(er repositories.EventsRepository) *eventsService {
	return &eventsService{repository: er}
}

func (es eventsService) GetAll() ([]models.Event, error) {
	events, err := es.repository.GetAll()
	if len(events) == 0 {
		err = errors.New("No events added yet")
	}

	return events, err
}

func (es eventsService) GetByID(id int64) (*models.Event, error) {
	return es.repository.GetByID(id)
}

func (es eventsService) Save(event *models.Event) error {
	return es.repository.Save(event)
}

func (es eventsService) Update(event *models.Event) error {
	return es.repository.Update(event)
}

func (es eventsService) Delete(event models.Event) error {
	return es.repository.Delete(event)
}
