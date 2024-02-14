package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"goproject.com/eventplanner-io/api/cmd/eventplannerd/dtos"
	"goproject.com/eventplanner-io/api/internal/models"
	"goproject.com/eventplanner-io/api/internal/platform/utils"
	"goproject.com/eventplanner-io/api/internal/services"
)

type usersHandler struct {
	service services.UsersService
}

func NewUsersHandler(service services.UsersService) *usersHandler {
	return &usersHandler{service}
}

func (uh usersHandler) Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = uh.service.Save(&user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (uh usersHandler) Login(context *gin.Context) {
	user, err := uh.parseUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = uh.service.ValidateCredentials(user)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func (uh usersHandler) parseUser(context *gin.Context) (*models.User, error) {
	var userData dtos.UserData
	err := context.ShouldBindJSON(&userData)

	if err != nil {
		return &models.User{}, err
	}

	user, err := uh.service.GetByEmail(userData.Email)
	user.Password = userData.Password

	return user, nil
}
