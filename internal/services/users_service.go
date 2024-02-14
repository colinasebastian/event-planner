package services

import (
	"goproject.com/eventplanner-io/api/internal/models"
	"goproject.com/eventplanner-io/api/internal/repositories"
)

type UsersService interface {
	GetByID(id int64) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Save(user *models.User) error
	ValidateCredentials(user *models.User) error
}

type usersService struct {
	repository repositories.UsersRepository
}

func NewUsersService(ur repositories.UsersRepository) *usersService {
	return &usersService{repository: ur}
}

func (us usersService) GetByID(id int64) (*models.User, error) {
	return us.repository.Get(&models.User{ID: id})
}

func (us usersService) GetByEmail(email string) (*models.User, error) {
	return us.repository.Get(&models.User{Email: email})
}

func (us usersService) Save(user *models.User) error {
	return us.repository.Save(user)
}

func (us usersService) ValidateCredentials(user *models.User) error {
	return us.repository.ValidateCredentials(user)
}
