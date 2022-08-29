package service

import (
	"goServerAuth/package/repository"
	"goServerAuth/structures"
)

type Autorization interface {
	CreateAdmin(user structures.User) (int, error)
	CreateUser(user structures.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
	GetUser(id int) (structures.User, error)
	EditUser(user structures.User) (error)
	DeleteUser(idUser int, idAdmin int) (error)
	GetUsers(id int) ([]structures.User, error)
}

type Service struct {
	Autorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
	}
}
