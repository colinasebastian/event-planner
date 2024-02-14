package services

import (
	"goproject.com/eventplanner-io/api/internal/models"
	"goproject.com/eventplanner-io/api/internal/repositories"
)

type RegistrationsService interface {
	Save(registration *models.Registration) error
	Delete(registration models.Registration) error
}

type registrationsService struct {
	repository repositories.RegistrationsRepository
}

func NewRegistrationsService(rr repositories.RegistrationsRepository) *registrationsService {
	return &registrationsService{repository: rr}
}

func (rs registrationsService) Save(registration *models.Registration) error {
	return rs.repository.Save(registration)
}

func (rs registrationsService) Delete(registration models.Registration) error {
	return rs.repository.Delete(registration)
}
