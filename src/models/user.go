package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	if err := user.format(stage); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("Name is mandatory and cannot be empty")
	}

	if user.Nick == "" {
		return errors.New("Nick is mandatory and cannot be empty")
	}

	if user.Email == "" {
		return errors.New("Email is mandatory and cannot be empty")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("Email: " + err.Error())
	}

	if stage == "create" && user.Password == "" {
		return errors.New("Password is mandatory and cannot be apty")
	}
	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "create" {
		passHash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(passHash)
	}
	return nil
}
