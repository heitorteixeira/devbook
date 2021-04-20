package repositories

import (
	"api/src/models"
	"database/sql"
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
