package study

import (
	"context"

	"github.com/joseluis244/db2dbmod/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type StudyStruct struct {
	client     *mongo.Client
	db         string
	collection *mongo.Collection
	ctx        context.Context
}

func New(client *mongo.Client, db string, collection string) *StudyStruct {
	return &StudyStruct{
		client:     client,
		db:         db,
		collection: client.Database(db).Collection(collection),
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

func (s *StudyStruct) FindToBuild() ([]struct {
	study     models.DestinationStudyType
	series    []models.DestinationSeriesType
	instances []models.DestinationInstanceType
}, error) {
	return nil, nil
}

func (s *StudyStruct) FindToSync() ([]models.DestinationStudyType, error) {
	return nil, nil
}
