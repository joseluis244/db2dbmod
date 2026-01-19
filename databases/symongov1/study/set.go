package study

import (
	"github.com/joseluis244/db2dbmod/databases/symongov1/models"
	"github.com/joseluis244/db2dbmod/databases/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//

func createStudy(StudyUuid string, Id int64, Tags map[string]interface{}) bson.M {
	return bson.M{
		"$setOnInsert": bson.M{
			"StudyUuid": StudyUuid,
			"Id":        Id,
		},
		"$set": bson.M{
			"Tags": Tags,
		},
	}
}

func (s *StudyStruct) UpsertStudies(studies []models.SyMongoV1StudyType) error {
	Models := []mongo.WriteModel{}
	for _, study := range studies {
		update := createStudy(study.StudyUuid, study.Id, study.Tags)
		filter := bson.M{"StudyUuid": study.StudyUuid}
		Model := mongo.NewUpdateOneModel()
		Model.SetFilter(filter)
		Model.SetUpdate(update)
		Model.SetUpsert(true)
		Models = append(Models, Model)
	}
	_, err := utils.BulkWrite(s.ctx, s.collection, Models)
	if err != nil {
		return err
	}
	return nil
}

func (s *StudyStruct) UpsertStudy(study models.SyMongoV1StudyType) error {
	opt := options.UpdateOne()
	opt.SetUpsert(true)
	filter := bson.M{"StudyUuid": study.StudyUuid}
	update := createStudy(study.StudyUuid, study.Id, study.Tags)
	_, err := s.collection.UpdateOne(s.ctx, filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}
