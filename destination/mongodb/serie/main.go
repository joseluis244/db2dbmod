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
	coll := client.Database(db).Collection(collection)
	coll.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.M{"StudyUuid": 1},
			Options: options.Index().SetUnique(false),
		},
		{
			Keys:    bson.M{"SerieUuid": 1},
			Options: options.Index().SetUnique(true),
		},
	})
	return &SerieStruct{
		client:     client,
		db:         db,
		collection: coll,
		ctx:        context.TODO(),
	}
}

func (s *SerieStruct) UpsertSeries(series []models.DestinationSeriesType) error {
	Models := []mongo.WriteModel{}
	for _, serie := range series {
		Model := mongo.NewUpdateOneModel().SetFilter(
			bson.M{"StudyUuid": serie.StudyUuid, "SerieUuid": serie.SerieUuid}).SetUpdate(
			bson.M{
				"$setOnInsert": bson.M{
					"DealerID":  serie.DealerID,
					"ClientID":  serie.ClientID,
					"BranchID":  serie.BranchID,
					"StudyUuid": serie.StudyUuid,
					"SerieUuid": serie.SerieUuid,
					"CreatedAt": serie.CreatedAt,
				},
				"$set": bson.M{
					"Tags":      serie.Tags,
					"UpdatedAt": serie.UpdatedAt,
					"Sync": models.SyncType{
						Status:   "pending",
						SyncTime: 0,
					},
				},
			}).SetUpsert(true)
		Models = append(Models, Model)
		if len(Models) == 500 {
			_, err := s.collection.BulkWrite(s.ctx, Models)
			if err != nil {
				return err
			}
			Models = []mongo.WriteModel{}
		}
	}
	if len(Models) > 0 {
		_, err := s.collection.BulkWrite(s.ctx, Models)
		if err != nil {
			return err
		}
	}
	return nil
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
			"CreatedAt": serie.CreatedAt,
		},
		"$set": bson.M{
			"Tags":      serie.Tags,
			"UpdatedAt": serie.UpdatedAt,
			"Sync": models.SyncType{
				Status:   "pending",
				SyncTime: 0,
			},
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
