package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"goproject.com/eventplanner-io/api/internal/models"
	"goproject.com/eventplanner-io/api/internal/services"
)

type registrationsHandler struct {
	service       services.RegistrationsService
	eventsService services.EventsService
	usersService  services.UsersService
}

func NewRegistrationsHandler(s services.RegistrationsService,
	es services.EventsService, us services.UsersService) *registrationsHandler {
	return &registrationsHandler{service: s, eventsService: es, usersService: us}
}

func (rh registrationsHandler) RegisterForEvent(context *gin.Context) {

	registration, err := rh.parseRegistration(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Colud not parse event id."})
		return
	}

	err = rh.service.Save(registration)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registration created!", "registration": registration})
}

func (rh registrationsHandler) CancelRegistration(context *gin.Context) {
	registration, err := rh.parseRegistration(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Colud not parse event id."})
		return
	}

	userID := context.GetInt64("userID")

	if registration.Event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = rh.service.Delete(*registration)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Colud not cancel registration."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Cancelled!"})
}

func (rh registrationsHandler) parseRegistration(context *gin.Context) (*models.Registration, error) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		return &models.Registration{}, err
	}

	user, err := rh.usersService.GetByID(userID)

	if err != nil {
		return &models.Registration{}, errors.New("Could not get user.")
	}

	event, err := rh.eventsService.GetByID(eventID)

	if err != nil {
		return &models.Registration{}, errors.New("Could not get event.")
	}

	registration := models.Registration{
		EventID: eventID,
		UserID:  userID,
		User:    *user,
		Event:   *event,
	}

	return &registration, nil
}
