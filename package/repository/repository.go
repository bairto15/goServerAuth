package repository

import (
	"goServerAuth/structures"

	"github.com/jmoiron/sqlx"
)

type Autorization interface {	
	CreateAdmin(user structures.User) (int, error)
	CreateUser(user structures.User) (int, error)
	Auth(login, password string) (structures.User, error)
	GetUser(id int) (structures.User, error)
	EditUser(user structures.User) (error)
	DeleteUser(idUser int, idAmin int) (error)
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