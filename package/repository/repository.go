package repository

import (
	"goServerAuth/structures"

	"github.com/jmoiron/sqlx"
)


type Autorization interface {
	CreateUser(user structures.User) (int, error)
	GetUser(login, password string) (structures.User, error)
	GetUsers(id int) ([]structures.User, error)
}

type Repository struct {
	Autorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}