package service

import (
	"goServerAuth/package/repository"
	"goServerAuth/structures"
)

type Autorization interface {
	CreateUser(user structures.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, error)
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
