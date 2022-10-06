package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents the user's model
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare call the methods to validate and format the data
func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New(`The field 'name' is mandatory`)
	}
	if user.Nick == "" {
		return errors.New(`The field 'nick' is mandatory`)
	}
	if user.Email == "" {
		return errors.New(`The field 'email' is mandatory`)
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New(`The email sent is invalid`)
	}
	if step == "REGISTRATION" && user.Password == "" {
		return errors.New(`The field 'password' is mandatory`)
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
