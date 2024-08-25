package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stm.Close()
	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stm.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials are invalid")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials are invalid")
	}

	return nil
}
