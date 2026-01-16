package mongodb

import (
	"context"
	"fmt"

	"github.com/joseluis244/db2dbmod/destination/mongodb/instance"
	"github.com/joseluis244/db2dbmod/destination/mongodb/serie"
	"github.com/joseluis244/db2dbmod/destination/mongodb/study"
	v3 "github.com/joseluis244/db2dbmod/destination/mongodb/v3"
	"github.com/joseluis244/db2dbmod/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDB struct {
	client   *mongo.Client
	db       string
	Study    *study.StudyStruct
	Series   *serie.SerieStruct
	Instance *instance.InstanceStruct
	V3       *v3.V3Struct
}

func New() *MongoDB {
	return &MongoDB{
		client:   nil,
		db:       "",
		Study:    nil,
		Series:   nil,
		Instance: nil,
		V3:       nil,
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
	m.Study = study.New(client, db, models.DestinationStudyCollection)
	m.Series = serie.New(client, db, models.DestinationSeriesRawCollection)
	m.Instance = instance.New(client, db, models.DestinationInstanceRawCollection)
	m.V3 = v3.New(client, db, models.DestinationV3Collection)
	fmt.Println("Connected to MongoDB")
	return nil
}

func (m *MongoDB) Disconnect() error {
	if m.client == nil {
		return nil
	}
	return m.client.Disconnect(context.TODO())
}
