package sqlite3

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite3Repo struct {
	DB *sql.DB
}

func newSqlite3Repo(db *sql.DB) *Sqlite3Repo {
	return &Sqlite3Repo{
		DB: db,
	}
}

func (db Sqlite3Repo) InitDB() {
	var err error
	db.DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to databes.")
	}

	// Controls how many open connections you can have to this database
	// If yoy need more than 10, those other request will have to wait until a connection is available again
	db.DB.SetMaxOpenConns(10)
	// How many connections you want to keep open if no one's using these connections at the moment
	db.DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	// CreateUsersTable()
	// CreateEventsTable()
	// CreateRegistrationTable()
}
