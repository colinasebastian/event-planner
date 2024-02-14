package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int64  `gorm:"type:integer;primaryKey;autoIncrement"`
	Email    string `gorm:"type:text;not null;uniqueIndex" binding:"required"`
	Password string `gorm:"type:text;not null" binding:"required"`
}

func (u User) Save() error {
	// query := `
	// INSERT INTO users(email, password)
	// VALUES (?, ?)`

	// stmt, err := repository.DB.Prepare(query)

	// if err != nil {
	// 	return err
	// }

	// defer stmt.Close()

	// hashedPassword, err := utils.HashPassword(u.Password)

	// if err != nil {
	// 	return err
	// }

	// result, err := stmt.Exec(u.Email, hashedPassword)

	// if err != nil {
	// 	return err
	// }

	// userID, err := result.LastInsertId()
	// u.ID = userID

	//return err
	return errors.New("test")
}

func (u *User) ValidateCredentials() error {
	// query := `
	// SELECT id, password
	// FROM users
	// WHERE email = ?
	// `

	// row := repository.DB.QueryRow(query, u.Email)

	// var retrievedPassword string
	// err := row.Scan(&u.ID, &retrievedPassword)

	// if err != nil {
	// 	return err
	// }

	// passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	// if !passwordIsValid {
	// 	return errors.New("Invalid credentials.")
	// }

	// return nil
	return errors.New("test")
}
