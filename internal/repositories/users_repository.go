package repositories

import (
	"errors"

	"goproject.com/eventplanner-io/api/internal/models"
	"goproject.com/eventplanner-io/api/internal/platform/utils"
	"gorm.io/gorm"
)

type UsersRepository interface {
	Get(params *models.User) (*models.User, error)
	Save(user *models.User) error
	ValidateCredentials(user *models.User) error
}

type usersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *usersRepository {
	return &usersRepository{
		db: db,
	}
}

func (ur *usersRepository) Get(params *models.User) (*models.User, error) {
	var user models.User
	err := ur.db.Model(&user).Where(params).FirstOrInit(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *usersRepository) Save(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}
	// Check if ID is updated
	user.Password = hashedPassword
	return ur.db.Create(&user).Error
}

func (ur *usersRepository) ValidateCredentials(user *models.User) error {
	var u models.User

	err := ur.db.Model(&u).Where("email = ?", &user.Email).FirstOrInit(&u).Error

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, u.Password)

	if !passwordIsValid {
		return errors.New("Invalid credentials.")
	}

	return nil
}
