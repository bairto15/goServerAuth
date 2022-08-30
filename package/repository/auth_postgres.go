package repository

import (
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
	query := "INSERT INTO admins (name, login, password, role) values ($1, $2, $3, $4) RETURNING id"

	row := r.db.QueryRow(query, user.Name, user.Login, user.Password, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

//Создание нового пользователя
func (r *AuthPostgres) CreateUser(user structures.User) (int, error) {
	var id int
	query := "INSERT INTO users (name, login, password, root, role) values ($1, $2, $3, $4, $5) RETURNING id"

	row := r.db.QueryRow(query, user.Name, user.Login, user.Password, user.Root, user.Role)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

//Авторизировать пользователя
func (r *AuthPostgres) Auth(login, password string) (structures.User, error) {
	var user structures.User

	queryUser := "SELECT * FROM users WHERE login=$1 AND password=$2"
	queryAdmin := "SELECT * FROM admins WHERE login=$1 AND password=$2"
	
	err := r.db.Get(&user, queryUser, login, password)
	if err != nil {
		err = r.db.Get(&user, queryAdmin, login, password)
	}

	return user, err
}
