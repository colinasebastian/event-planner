package repositories_tests

import (
	"os"
	"os/exec"
	"testing"
	"time"

	"goproject.com/eventplanner-io/api/internal/models"
	"goproject.com/eventplanner-io/api/internal/platform/database/migrations"
	"goproject.com/eventplanner-io/api/internal/platform/database/sqlite"
	"gorm.io/gorm"
)

var gdb *gorm.DB

func TestMain(m *testing.M) {
	var err error
	dbName := "test_database.db"
	exec.Command("rm", "-f", dbName)

	gdb, err = sqlite.Connect(dbName)
	if err != nil {
		panic(("Failed to connect to database"))
	}

	err = migrations.MigrateDB(gdb)
	if err != nil {
		panic(("Failed to use migrations."))
	}

	date, _ := time.Parse(time.RFC3339, "2024-01-01T12:00:00.000Z")

	users := []models.User{
		{
			ID:       1,
			Email:    "test1@example.com",
			Password: "$2a$14$0Vh7mjHrEWLiS544KW2r.uoVWBpwWEV8foJxhOIj2plLcF9HfHzXq",
		}, {
			ID:       2,
			Email:    "test2@example.com",
			Password: "$2a$14$0Vh7mjHrEWLiS544KW2r.uoVWBpwWEV8foJxhOIj2plLcF9HfHzXq",
		}, {
			ID:       3,
			Email:    "test3@example.com",
			Password: "$2a$14$0Vh7mjHrEWLiS544KW2r.uoVWBpwWEV8foJxhOIj2plLcF9HfHzXq",
		},
	}

	events := []models.Event{
		{
			ID:          1,
			Name:        "Test event",
			Description: "A test event",
			Location:    "A test location",
			DateTime:    date,
			UserID:      users[0].ID,
			User:        users[0],
		}, {
			ID:          2,
			Name:        "Test event 2",
			Description: "A test event 2",
			Location:    "A test location 2",
			DateTime:    date,
			UserID:      users[0].ID,
			User:        users[0],
		},
	}

	registrations := []models.Registration{
		{
			EventID: events[0].ID,
			UserID:  users[0].ID,
			User:    users[0],
			Event:   events[0],
		},
		{
			EventID: events[1].ID,
			UserID:  users[0].ID,
			User:    users[0],
			Event:   events[1],
		},
	}

	gdb.Create(users)
	gdb.Create(events)
	gdb.Create(registrations)
	os.Exit(m.Run())
}
