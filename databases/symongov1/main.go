package symongov2

import (
	"context"
	"fmt"

	"github.com/joseluis244/db2dbmod/databases/symongov2/instance"
	"github.com/joseluis244/db2dbmod/databases/symongov2/serie"
	"github.com/joseluis244/db2dbmod/databases/symongov2/study"
	"github.com/joseluis244/db2dbmod/databases/symongov2/system"
	v3 "github.com/joseluis244/db2dbmod/databases/symongov2/v3"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type SyMongoV2 struct {
	client   *mongo.Client
	db       string
	Study    *study.StudyStruct
	Series   *serie.SerieStruct
	Instance *instance.InstanceStruct
	V3       *v3.V3Struct
	System   *system.SystemStruct
}

func New() *SyMongoV2 {
	return &SyMongoV2{
		client:   nil,
		db:       "",
		Study:    nil,
		Series:   nil,
		Instance: nil,
		V3:       nil,
		System:   nil,
	}
}

func (m *SyMongoV2) Connect(dsn string, db string) error {
	client, err := mongo.Connect(options.Client().ApplyURI(dsn))
	if err != nil {
		return err
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return err
	}
	m.client = client
	m.db = db
	m.Study = study.New(client, db, "StudyRaw")
	m.Series = serie.New(client, db, "SerieRaw")
	m.Instance = instance.New(client, db, "InstanceRaw")
	m.V3 = v3.New(client, db, "V3")
	m.System = system.New(client, db, "System")
	fmt.Println("Connected to MongoDB")
	return nil
}

func (m *SyMongoV2) Disconnect() error {
	if m.client == nil {
		return nil
	}
	return m.client.Disconnect(context.TODO())
}
