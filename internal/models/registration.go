package models

import (
	"errors"

	"gorm.io/gorm"
)

type Registration struct {
	gorm.Model
	ID      int64 `gorm:"type:integer;primaryKey;autoIncrement"`
	EventID int64 `gorm:"type:integer;not null" binding:"required"`
	UserID  int64 `gorm:"type:integer;not null" binding:"required"`
	User    User  `gorm:"foreignKey:UserID"`
	Event   Event `gorm:"foreignKey:EventID"`
}

func (r Registration) Save() error {
	// query := `
	// INSERT INTO registrations(event_id, user_id)
	// VALUES (?, ?)`
	// stmt, err := repository.DB.Prepare(query)

	// if err != nil {
	// 	return err
	// }

	// defer stmt.Close()

	// result, err := stmt.Exec(r.EventID, r.UserID)

	// if err != nil {
	// 	return err
	// }

	// id, err := result.LastInsertId()
	// r.ID = id

	// return err

	return errors.New("test")
}

func (r Registration) Delete() error {
	// query := `DELETE FROM registrations WHERE event_id = ? AND user_id = ?`
	// stmt, err := repository.DB.Prepare(query)

	// if err != nil {
	// 	return err
	// }

	// defer stmt.Close()

	// _, err = stmt.Exec(r.EventID, r.UserID)

	// return nil

	return errors.New("test")
}
