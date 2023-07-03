package repository

import (
	"database/sql"
	"forum/entity"
)

func (r repo) CreateUser(user *entity.User) error {
	stmt, err := r.db.Prepare(`INSERT INTO User (Username, Email, Password) VALUES (?, ?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// CheckUsername checks if a username exists in the User table
func (r repo) CheckUsername(username string) (bool, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM User WHERE Username = ?;`, username).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// CheckEmail checks if an email exists in the User table
func (r repo) CheckEmail(email string) (bool, error) {
	var count int
	err := r.db.QueryRow(`SELECT COUNT(*) FROM User WHERE Email = ?;`, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// GetUserByEmail retrieves a user by email from the User table
func (r repo) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.QueryRow(`SELECT ID, Username, Email, Password FROM User WHERE Email = ?;`, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // User not found
		}
		return nil, err
	}

	return &user, nil
}
