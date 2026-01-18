package study

import (
	"context"

	"github.com/joseluis244/db2dbmod/databases/symongov2/models"
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

func (s *StudyStruct) FindByStudyUuid(studyUuid string) (models.DestinationStudyType, error) {
	var study models.DestinationStudyType
	err := s.collection.FindOne(s.ctx, bson.M{
		"StudyUuid": studyUuid,
	}).Decode(&study)
	if err != nil {
		return models.DestinationStudyType{}, err
	}
	return study, nil
}

func (s *StudyStruct) GetToBuild(filter bson.M) ([]struct {
	Study     models.DestinationStudyType
	Series    []models.DestinationSeriesType
	Instances []models.DestinationInstanceType
}, error) {
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
		models.DestinationStudyType `bson:",inline"`
		Series                      []models.DestinationSeriesType   `bson:"Series"`
		Instances                   []models.DestinationInstanceType `bson:"Instances"`
	}

	// Verificar error del cursor.All
	if err := cursor.All(s.ctx, &responses); err != nil {
		return nil, err
	}

	// Simplificar: retornar directamente sin conversión innecesaria
	var res []struct {
		Study     models.DestinationStudyType
		Series    []models.DestinationSeriesType
		Instances []models.DestinationInstanceType
	}

	for _, response := range responses {
		res = append(res, struct {
			Study     models.DestinationStudyType
			Series    []models.DestinationSeriesType
			Instances []models.DestinationInstanceType
		}{
			Study:     response.DestinationStudyType,
			Series:    response.Series,
			Instances: response.Instances,
		})
	}

	return res, nil
}

func (s *StudyStruct) GetToSync() ([]models.DestinationStudyType, error) {
	return nil, nil
}
