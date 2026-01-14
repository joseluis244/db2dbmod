package serie

import (
	"context"

	"github.com/joseluis244/db2dbmod/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type SerieStruct struct {
	client     *mongo.Client
	db         string
	collection *mongo.Collection
	ctx        context.Context
}

func New(client *mongo.Client, db string, collection string) *SerieStruct {
	return &SerieStruct{
		client:     client,
		db:         db,
		collection: client.Database(db).Collection(collection),
		ctx:        context.TODO(),
	}
}

func (s *SerieStruct) UpsertSerie(serie models.DestinationSeriesType) error {
	opt := options.UpdateOne()
	opt.SetUpsert(true)

	_, err := s.collection.UpdateOne(s.ctx, bson.M{
		"StudyUuid": serie.StudyUuid,
		"SerieUuid": serie.SerieUuid,
	}, bson.M{
		"$setOnInsert": bson.M{
			"DealerID":  serie.DealerID,
			"ClientID":  serie.ClientID,
			"BranchID":  serie.BranchID,
			"StudyUuid": serie.StudyUuid,
			"SerieUuid": serie.SerieUuid,
			"Tags":      serie.Tags,
			"CreatedAt": serie.CreatedAt,
			"BuildTime": 0,
			"Sync": models.SyncType{
				Status:   "pending",
				SyncTime: 0,
			},
		},
		"$set": bson.M{
			"Tags":      serie.Tags,
			"UpdatedAt": serie.UpdatedAt,
		},
	},
		opt)
	if err != nil {
		return err
	}
	return nil
}

func (s *SerieStruct) SetUpdatedAt(serieUuid string, updatedAt int64) error {
	_, err := s.collection.UpdateOne(s.ctx, bson.M{
		"SerieUuid": serieUuid,
	}, bson.M{
		"$set": bson.M{
			"UpdatedAt": updatedAt,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *SerieStruct) SetSync(serieUuid string, status string, syncTime int64) error {
	_, err := s.collection.UpdateOne(s.ctx, bson.M{
		"SerieUuid": serieUuid,
	}, bson.M{
		"$set": bson.M{
			"Sync": models.SyncType{
				Status:   status,
				SyncTime: syncTime,
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *SerieStruct) FindByStudyUuid(studyUuid string) (models.DestinationSeriesType, error) {
	var series models.DestinationSeriesType
	err := s.collection.FindOne(s.ctx, bson.M{
		"StudyUuid": studyUuid,
	}).Decode(&series)
	if err != nil {
		return models.DestinationSeriesType{}, err
	}
	return series, nil
}
