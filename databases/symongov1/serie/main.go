package serie

import (
	"context"

	"github.com/joseluis244/db2dbmod/databases/symongov1/models"
	"github.com/joseluis244/db2dbmod/databases/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func createSerie(SerieUuid string, Id int64, StudyUuid string, Tags map[string]interface{}) bson.M {
	return bson.M{
		"$setOnInsert": bson.M{
			"SerieUuid": SerieUuid,
			"Id":        Id,
			"StudyUuid": StudyUuid,
		},
		"$set": bson.M{
			"tags": Tags,
		},
	}
}

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
			Options: options.Index().SetUnique(false),
		},
	})
	return &SerieStruct{
		client:     client,
		db:         db,
		collection: coll,
		ctx:        context.TODO(),
	}
}

func (s *SerieStruct) UpsertSeries(series []models.SyMongoV1SeriesType) error {
	Models := []mongo.WriteModel{}
	for _, serie := range series {
		filter := bson.M{"StudyUuid": serie.StudyUuid, "SerieUuid": serie.SerieUuid}
		update := createSerie(serie.SerieUuid, serie.Id, serie.StudyUuid, serie.Tags)
		Model := mongo.NewUpdateOneModel()
		Model.SetFilter(filter)
		Model.SetUpdate(update).SetUpsert(true)
		Models = append(Models, Model)
	}
	_, err := utils.BulkWrite(s.ctx, s.collection, Models)
	if err != nil {
		return err
	}
	return nil
}

func (s *SerieStruct) UpsertSerie(serie models.SyMongoV1SeriesType) error {
	opt := options.UpdateOne()
	opt.SetUpsert(true)
	filter := bson.M{"StudyUuid": serie.StudyUuid, "SerieUuid": serie.SerieUuid}
	update := createSerie(serie.SerieUuid, serie.Id, serie.StudyUuid, serie.Tags)
	_, err := s.collection.UpdateOne(s.ctx, filter, update, opt)
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

func (s *SerieStruct) FindByStudyUuid(studyUuid string) (models.SyMongoV1SeriesType, error) {
	var series models.SyMongoV1SeriesType
	err := s.collection.FindOne(s.ctx, bson.M{
		"StudyUuid": studyUuid,
	}).Decode(&series)
	if err != nil {
		return models.SyMongoV1SeriesType{}, err
	}
	return series, nil
}
