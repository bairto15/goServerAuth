package service

import "goServerAuth/package/repository"

type SendContent interface {
}

type Service struct {
	SendContent
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
