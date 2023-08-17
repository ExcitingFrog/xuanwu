package mongodb

import (
	"context"

	"github.com/ExcitingFrog/xuanwu/pkg/provider"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	provider.IProvider
	Config *Config

	Client *mongo.Client
}

func NewMongoDB(config *Config) *MongoDB {
	if config == nil {
		config = NewConfig()
	}
	return &MongoDB{
		Config: config,
	}
}

func (p *MongoDB) Run() error {
	opts := options.Client().
		ApplyURI(p.Config.URI).
		SetConnectTimeout(p.Config.Timeout).
		SetMaxPoolSize(p.Config.MaxPoolSize).
		SetMaxConnIdleTime(p.Config.MaxIdle)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	p.Client = client

	return nil
}

func (m *MongoDB) Close() error {
	if err := m.Client.Disconnect(context.Background()); err != nil {
		return err
	}

	return nil
}
