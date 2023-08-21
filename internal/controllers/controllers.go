package controllers

import (
	"github.com/ExcitingFrog/xuanwu/internal/services"
)

type Controllers struct {
	service services.IService
}

func NewControllers(service services.IService) *Controllers {
	return &Controllers{service: service}
}
