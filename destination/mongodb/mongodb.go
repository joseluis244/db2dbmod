package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDB struct {
	client *mongo.Client
	db     string
}

func New() *MongoDB {
	return &MongoDB{
		client: nil,
		db:     "",
	}
}

func (m *MongoDB) Connect(dsn string, db string) error {
	client, err := mongo.Connect(options.Client().ApplyURI(dsn))
	if err != nil {
		return err
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return err
	}
	m.client = client
	m.db = db
	return nil
}

func (m *MongoDB) Disconnect() error {
	if m.client == nil {
		return nil
	}
	return m.client.Disconnect(context.TODO())
}
