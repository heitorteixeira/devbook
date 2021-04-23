package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

//NewUserRepository create one user repository
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

//Create insert one user in DB
func (userRepository UserRepository) Create(user models.User) (uint64, error) {
	statement, err := userRepository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertId), nil
}

//Find return all users by name or nick
func (userRepository UserRepository) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := userRepository.db.Query(
		"select id, name, nick, email, criatedAt From users where name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CriatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
