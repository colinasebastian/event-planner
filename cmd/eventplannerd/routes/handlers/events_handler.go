package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"goproject.com/eventplanner-io/api/cmd/eventplannerd/dtos"
	"goproject.com/eventplanner-io/api/internal/models"
	"goproject.com/eventplanner-io/api/internal/services"
)

type eventsHandler struct {
	service     services.EventsService
	userService services.UsersService
}

func NewEventsHandler(s services.EventsService, us services.UsersService) *eventsHandler {
	return &eventsHandler{service: s, userService: us}
}

func (eh eventsHandler) GetEvents(context *gin.Context) {
	events, err := eh.service.GetAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later.", "error": err})
		return
	}
	context.JSON(http.StatusOK, events)
}

func (eh eventsHandler) GetEventByID(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Colud not parse event id."})
		return
	}
	event, err := eh.service.GetByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Colud not fetch event."})
		return
	}

	context.JSON(http.StatusOK, event)
}

func (eh eventsHandler) CreateEvent(context *gin.Context) {

	event, err := eh.parseEvent(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = eh.service.Save(event)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func (eh eventsHandler) UpdateEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Colud not parse event id."})
		return
	}

	userID := context.GetInt64("userID")
	event, err := eh.service.GetByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Colud not fetch event."})
		return
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}

	updatedEvent, err := eh.parseEvent(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	updatedEvent.ID = eventID
	err = eh.service.Update(updatedEvent)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Colud not update event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
}

func (eh eventsHandler) DeleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Colud not parse event id."})
		return
	}

	userID := context.GetInt64("userID")
	event, err := eh.service.GetByID(eventID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Colud not fetch event."})
		return
	}

	if event.UserID != userID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = eh.service.Delete(*event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Colud not delete event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}

func (eh eventsHandler) parseEvent(context *gin.Context) (*models.Event, error) {
	var eventData dtos.EventData
	err := context.ShouldBindJSON(&eventData)

	if err != nil {
		return &models.Event{}, err
	}

	userID := context.GetInt64("userID")
	user, err := eh.userService.GetByID(userID)

	event := models.Event{
		Name:        eventData.Name,
		Description: eventData.Description,
		Location:    eventData.Location,
		DateTime:    eventData.DateTime,
		User:        *user,
	}

	return &event, nil
}
