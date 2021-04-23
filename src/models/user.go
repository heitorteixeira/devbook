package models

import (
	"errors"
	"strings"
	"time"
)

//User represent one user using social media
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CriatedAt time.Time `json:"criatedAt,omitempty"`
}

//Prepare will call methods to validate and format User
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("Name is mandatory and cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("Nick is mandatory and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("Email is mandatory and cannot be empty")
	}

	if user.Password == "" {
		return errors.New("Password is mandatory and cannot be apty")
	}
	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
