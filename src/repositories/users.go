package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

// Users represents a user repository
type Users struct {
	db *sql.DB
}

// NewUserRepository create a users repository
func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

// Create do an INSERT in database
func (respository Users) Create(user models.User) (uint64, error) {
	statement, err := respository.db.Prepare("INSERT INTO users (name, nick, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

// Find do a GET in database
func (respository Users) Find(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%

	data, err := respository.db.Query(
		`SELECT id, name, nick, email, createdAt FROM users WHERE name LIKE ? OR nick LIKE ?`,
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	var users []models.User

	for data.Next() {
		var user models.User

		if err = data.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// FindByID is a method to GET an unique user by ID
func (respository Users) FindByID(ID uint64) (models.User, error) {
	data, err := respository.db.Query(
		`SELECT id, name, nick, email, createdAt FROM users WHERE ID = ?`,
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer data.Close()

	var user models.User

	if data.Next() {
		if err = data.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// UpdateUser is method to update an unique user by ID
func (respository Users) UpdateUser(ID uint64, user models.User) error {
	statement, err := respository.db.Prepare(
		"UPDATE users SET name = ?, nick = ?, email = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// DeleteUser is method to REMOVE an user by ID
func (respository Users) DeleteUser(ID uint64) error {
	statement, err := respository.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// FindByEmail is a method to GET user by email
func (repository Users) FindByEmail(email string) (models.User, error) {
	data, err := repository.db.Query("SELECT id, password FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer data.Close()

	var user models.User

	if data.Next() {
		if err = data.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Follow receive an user and a follower
func (repository Users) Follow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow receive an user to unfollow from other
func (repository Users) Unfollow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"DELETE FROM followers WHERE user_id = ? AND follower_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

func (repository Users) FindFollowers(userID uint64) ([]models.User, error) {
	rows, err := repository.db.Query(`
	SELECT u.id, u.name, u.nick, u.email, u.createdAt
	FROM users u INNER JOIN followers f ON u.id = f.follower_id
	WHERE f.user_id = ?
	`, userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []models.User
	for rows.Next() {
		var follower models.User

		if err = rows.Scan(
			&follower.ID,
			&follower.Name,
			&follower.Nick,
			&follower.Email,
			&follower.CreatedAt,
		); err != nil {
			return nil, err
		}

		followers = append(followers, follower)
	}

	return followers, nil
}
