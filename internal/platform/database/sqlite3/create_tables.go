package sqlite3

// func CreateUsersTable() {
// 	createUsersTable := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		email TEXT NOT NULL UNIQUE,
// 		password TEXT NOT NULL
// 	)
// 	`
// 	_, err := DB.Exec(createUsersTable)

// 	if err != nil {
// 		panic("Could not create users table.")
// 	}
// }

// func CreateEventsTable() {
// 	createEventsTable := `
// 	CREATE TABLE IF NOT EXISTS events (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		name TEXT NOT NULL,
// 		description TEXT NOT NULL,
// 		location TEXT NOT NULL,
// 		dateTime DATETIME NOT NULL,
// 		user_id INTEGER,
// 		FOREIGN KEY(user_id) REFERENCES users(id)
// 	)
// 	`

// 	_, err := DB.Exec(createEventsTable)

// 	if err != nil {
// 		panic("Could not create events table.")
// 	}
// }

// func CreateRegistrationTable() {
// 	createRegistrationTable := `
// 	CREATE TABLE IF NOT EXISTS registrations (
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		event_id INTEGER,
// 		user_id INTEGER,
// 		FOREIGN KEY (event_id) REFERENCES events (id),
// 		FOREIGN KEY (user_id) REFERENCES users (id)
// 	)
// 	`

// 	_, err := DB.Exec(createRegistrationTable)

// 	if err != nil {
// 		panic("Could not create registrations table.")
// 	}
// }
