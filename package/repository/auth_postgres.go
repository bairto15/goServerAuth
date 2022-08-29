package repository

import (
	"fmt"
	"goServerAuth/structures"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

//Создание нового админа
func (r *AuthPostgres) CreateAdmin(user structures.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, login, password) values ($1, $2, $3) RETURNING id", "admins")

	row := r.db.QueryRow(query, user.Name, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

//Создание нового пользователя
func (r *AuthPostgres) CreateUser(user structures.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, login, password, root) values ($1, $2, $3, $4) RETURNING id", "users")

	row := r.db.QueryRow(query, user.Name, user.Login, user.Password, user.Root)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

//Авторизировать пользователя
func (r *AuthPostgres) Auth(login, password string) (structures.User, error) {
	var user structures.User

	query := "SELECT admins.id, users.id FROM users JOIN admins ON admins.id = users.root WHERE (users.login=$1 AND users.password=$2) OR (admins.login=$1 AND admins.password=$2)"
	err := r.db.Get(&user, query, login, password)

	return user, err
}
