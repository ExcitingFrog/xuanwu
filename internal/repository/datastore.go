package repository

import "github.com/ExcitingFrog/go-core-common/mongodb"

type IRepository interface {
	IHello
}

type repository struct {
	mongo *mongodb.MongoDB
}

func NewRepository(mongo *mongodb.MongoDB) IRepository {
	return &repository{
		mongo: mongo,
	}
}
