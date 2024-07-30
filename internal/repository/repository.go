package repository

import (
	"database/sql"
	"github.com/j0n4t45d3v/crud/internal/entity"
)

type IUserRepository interface {
	FindAll() ([]entity.User, error)
	Save(user entity.User) (int, error)
	Delete(id string) error
	Update(userUpdate entity.User, idUser string) error
}

type Repository struct {
	con *sql.DB
}

func NewRepository(connection *sql.DB) *Repository {
	return &Repository{con: connection}
}

func (r Repository) FindAll() ([]entity.User, error) {

	users := []entity.User{}

	query := "SELECT u.username, u.email, u.password FROM users u"
	rows, err := r.con.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user entity.User
		rows.Scan(&user.Name, &user.Email, &user.Password)
		users = append(users, user)
	}

	if errorFetch := rows.Err(); errorFetch != nil {
		return nil, err
	}

	return users, nil
}

func (r Repository) Save(user entity.User) (int, error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := r.con.Exec(query, user.Name, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}
	return int(userId), nil
}

func (r Repository) Delete(id string) error {

	query := "DELETE FROM users WHERE id = ?"

	_, err := r.con.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r Repository) Update(userUpdate entity.User, idUser string) error {
	query := "UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?"

	_, err := r.con.Exec(
		query,
		userUpdate.Name,
		userUpdate.Email,
		userUpdate.Password,
		idUser,
	)

	if err != nil {
		return err
	}

	return nil
}
