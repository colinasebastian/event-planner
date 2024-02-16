package repositories_tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"goproject.com/eventplanner-io/api/internal/models"
	"goproject.com/eventplanner-io/api/internal/repositories"
	"gorm.io/gorm"
)

func TestEventGetAll_Success(t *testing.T) {
	er := repositories.NewEventsRepository(gdb)

	events, err := er.GetAll()

	assert.Nil(t, err)
	assert.Len(t, events, len(events))
	assert.Equal(t, nil, err)
}

func TestEventGetByID_Success(t *testing.T) {
	er := repositories.NewEventsRepository(gdb)

	eventToCheck := models.Event{
		ID:          2,
		Name:        "Test event 2",
		Description: "A test event 2",
		Location:    "A test location 2",
		DateTime:    time.Now(),
		UserID:      1,
		User: models.User{
			ID:       1,
			Email:    "test1@example.com",
			Password: "$2a$14$0Vh7mjHrEWLiS544KW2r.uoVWBpwWEV8foJxhOIj2plLcF9HfHzXq",
		},
	}

	event, err := er.GetByID(2)

	assert.Nil(t, err)
	assert.Equal(t, event.ID, eventToCheck.ID)
	assert.Equal(t, event.Name, eventToCheck.Name)
	assert.Equal(t, event.Description, eventToCheck.Description)
	assert.Equal(t, event.Location, eventToCheck.Location)
	assert.Equal(t, nil, err)
}

func TestEventGetByID_Fail(t *testing.T) {
	er := repositories.NewEventsRepository(gdb)

	eventToCheck := models.Event{
		ID:          2,
		Name:        "Test event 2",
		Description: "A test event 2",
		Location:    "A test location 2",
		DateTime:    time.Now(),
		UserID:      1,
		User: models.User{
			ID:       1,
			Email:    "test1@example.com",
			Password: "$2a$14$0Vh7mjHrEWLiS544KW2r.uoVWBpwWEV8foJxhOIj2plLcF9HfHzXq",
		},
	}

	event, err := er.GetByID(1)
	assert.Nil(t, err)

	assert.NotEqual(t, event.ID, eventToCheck.ID)
	assert.NotEqual(t, event.Name, eventToCheck.Name)
	assert.NotEqual(t, event.Description, eventToCheck.Description)
	assert.NotEqual(t, event.Location, eventToCheck.Location)
	assert.Equal(t, nil, err)
}

func TestEventGetByID_Fail_NotFound(t *testing.T) {
	er := repositories.NewEventsRepository(gdb)

	event, err := er.GetByID(1000000000)
	fmt.Println(event)
	assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
}

func TestEventCreate_Success(t *testing.T) {
	er := repositories.NewEventsRepository(gdb)
	date, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00.000Z")

	eventParam := models.Event{
		Name:        "Test event 3",
		Description: "A test event 3",
		Location:    "A test location 3",
		DateTime:    date,
	}

	err := er.Save(&eventParam)
	assert.Nil(t, err)

	event, err := er.GetByID(3)

	assert.Equal(t, event.Name, eventParam.Name)
	assert.Equal(t, nil, err)
}

func TestEventUpdate_Success(t *testing.T) {
	er := repositories.NewEventsRepository(gdb)
	date, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00.000Z")

	eventParam := models.Event{
		ID:          3,
		Name:        "Test event 3 updated",
		Description: "A test event 3 updated",
		Location:    "A test location 3 updated",
		DateTime:    date,
		UserID:      1,
		User: models.User{
			ID:       1,
			Email:    "test1@example.com",
			Password: "$2a$14$0Vh7mjHrEWLiS544KW2r.uoVWBpwWEV8foJxhOIj2plLcF9HfHzXq",
		},
	}

	err := er.Update(&eventParam)
	assert.Nil(t, err)

	event, err := er.GetByID(3)

	assert.Equal(t, event.Name, eventParam.Name)
	assert.Equal(t, nil, err)
}

func TestEventDeleted_Success(t *testing.T) {
	er := repositories.NewEventsRepository(gdb)
	date, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00.000Z")

	eventParam := models.Event{
		ID:          3,
		Name:        "Test event 3 updated",
		Description: "A test event 3 updated",
		Location:    "A test location 3 updated",
		DateTime:    date,
		UserID:      1,
		User: models.User{
			ID:       1,
			Email:    "test1@example.com",
			Password: "$2a$14$0Vh7mjHrEWLiS544KW2r.uoVWBpwWEV8foJxhOIj2plLcF9HfHzXq",
		},
	}

	err := er.Delete(eventParam)
	assert.Nil(t, err)

	_, err = er.GetByID(3)

	assert.Equal(t, nil, err)
}
