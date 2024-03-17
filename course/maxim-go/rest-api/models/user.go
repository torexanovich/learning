package models

import (
	"errors"
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPass)
	if err != nil {
		return err
	}

	userID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userID
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var gotPass string

	err := row.Scan(&u.ID, &gotPass)
	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, gotPass)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}
