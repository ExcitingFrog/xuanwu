package services

import (
	"github.com/ExcitingFrog/xuanwu/internal/repository"
)

type IService interface {
	IHello
}

type Service struct {
	repository repository.IRepository
}

func NewService(repository repository.IRepository) IService {
	return &Service{
		repository: repository,
	}
}
