package services

import (
	"github.com/ExcitingFrog/xuanwu/internal/repository"
	"github.com/ExcitingFrog/xuanwu/internal/resources"
)

type IService interface {
	IHello
}

type Service struct {
	repository repository.IRepository
	xuyu       *resources.Xuyu
}

func NewService(repository repository.IRepository, xuyu *resources.Xuyu) IService {
	return &Service{
		repository: repository,
		xuyu:       xuyu,
	}
}
