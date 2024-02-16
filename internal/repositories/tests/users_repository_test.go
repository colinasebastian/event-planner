package repositories_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"goproject.com/eventplanner-io/api/internal/models"
	"goproject.com/eventplanner-io/api/internal/repositories"
)

func TestUserCreate_Success(t *testing.T) {
	ur := repositories.NewUsersRepository(gdb)

	userParam := models.User{
		Email:    "test4@example.com",
		Password: "test",
	}

	err := ur.Save(&userParam)
	assert.Nil(t, err)

	user, err := ur.Get(&models.User{Email: userParam.Email})

	assert.Equal(t, user.Email, userParam.Email)
	assert.Equal(t, nil, err)
}

func TestUserCreate_Fail_EmptyUser(t *testing.T) {
	ur := repositories.NewUsersRepository(gdb)

	userParam := models.User{}

	err := ur.Save(&userParam)
	assert.Nil(t, err)

	err = ur.Save(&userParam)
	assert.NotNil(t, err)
}

func TestValidateCredentials_Success(t *testing.T) {
	ur := repositories.NewUsersRepository(gdb)

	userParam := models.User{
		Email:    "test3@example.com",
		Password: "test",
	}

	err := ur.ValidateCredentials(&userParam)
	assert.Nil(t, err)
}

func TestValidateCredentials_Fail(t *testing.T) {
	ur := repositories.NewUsersRepository(gdb)

	userParam := models.User{
		Email:    "test3@example.com",
		Password: "test123",
	}

	err := ur.ValidateCredentials(&userParam)
	assert.NotNil(t, err)
	assert.Error(t, err, "Invalid credentials.")
}
