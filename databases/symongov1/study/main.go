package study

import (
	"context"

	"github.com/joseluis244/db2dbmod/databases/symongov1/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type StudyStruct struct {
	client     *mongo.Client
	db         string
	collection *mongo.Collection
	ctx        context.Context
}

func New(client *mongo.Client, db string, collection string) *StudyStruct {
	coll := client.Database(db).Collection(collection)
	coll.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.M{"StudyUuid": 1},
			Options: options.Index().SetUnique(false),
		},
		{
			Keys:    bson.M{"Tags.0008,0020": 1},
			Options: options.Index().SetUnique(false),
		},
	})
	return &StudyStruct{
		client:     client,
		db:         db,
		collection: coll,
		ctx:        context.TODO(),
	}
}

func (s *StudyStruct) FindByStudyUuid(studyUuid string) (models.SyMongoV1StudyType, error) {
	var study models.SyMongoV1StudyType
	err := s.collection.FindOne(s.ctx, bson.M{
		"StudyUuid": studyUuid,
	}).Decode(&study)
	if err != nil {
		return models.SyMongoV1StudyType{}, err
	}
	return study, nil
}

func (s *StudyStruct) GetToBuild(filter bson.M) ([]models.SyMongoV1ToBuildStruct, error) {
	Aggr := bson.A{
		bson.M{"$match": filter},
		bson.M{
			"$lookup": bson.M{
				"from":         "SeriesRaw",
				"localField":   "StudyUuid",
				"foreignField": "StudyUuid",
				"as":           "Series",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "InstanceRaw",
				"localField":   "StudyUuid",
				"foreignField": "StudyUuid",
				"as":           "Instances",
			},
		},
	}

	AggregationOptions := options.Aggregate()
	AggregationOptions.SetAllowDiskUse(true)

	cursor, err := s.collection.Aggregate(s.ctx, Aggr, AggregationOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(s.ctx) // IMPORTANTE: cerrar el cursor

	var responses []struct {
		models.SyMongoV1StudyType `bson:",inline"`
		Series                    []models.SyMongoV1SeriesType   `bson:"Series"`
		Instances                 []models.SyMongoV1InstanceType `bson:"Instances"`
	}

	// Verificar error del cursor.All
	if err := cursor.All(s.ctx, &responses); err != nil {
		return nil, err
	}

	// Simplificar: retornar directamente sin conversión innecesaria
	var res []models.SyMongoV1ToBuildStruct

	for _, response := range responses {
		res = append(res, models.SyMongoV1ToBuildStruct{
			Study:     response.SyMongoV1StudyType,
			Series:    response.Series,
			Instances: response.Instances,
		})
	}

	return res, nil
}
