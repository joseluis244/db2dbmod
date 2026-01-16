package study

import (
	"github.com/joseluis244/db2dbmod/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//

func createStudy(DealerID string, ClientID string, BranchID string, StudyUuid string, CreatedAt int64, Tags map[string]interface{}) bson.M {
	return bson.M{
		"$setOnInsert": bson.M{
			"DealerID":  DealerID,
			"ClientID":  ClientID,
			"BranchID":  BranchID,
			"StudyUuid": StudyUuid,
			"CreatedAt": CreatedAt,
			"BuildTime": 0,
			"Sync": models.SyncType{
				Status:   "pending",
				SyncTime: 0,
			},
		},
		"$set": bson.M{
			"Tags":      Tags,
			"UpdatedAt": 0,
		},
	}
}

func (s *StudyStruct) UpsertStudies(studies []models.DestinationStudyType) error {
	Models := []mongo.WriteModel{}
	for _, study := range studies {
		update := createStudy(study.DealerID, study.ClientID, study.BranchID, study.StudyUuid, study.CreatedAt, study.Tags)
		filter := bson.M{"StudyUuid": study.StudyUuid}
		Model := mongo.NewUpdateOneModel()
		Model.SetFilter(filter)
		Model.SetUpdate(update)
		Model.SetUpsert(true)
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

func (s *StudyStruct) UpsertStudy(study models.DestinationStudyType) error {
	opt := options.UpdateOne()
	opt.SetUpsert(true)
	filter := bson.M{"StudyUuid": study.StudyUuid}
	update := createStudy(study.DealerID, study.ClientID, study.BranchID, study.StudyUuid, study.CreatedAt, study.Tags)
	_, err := s.collection.UpdateOne(s.ctx, filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}

func (s *StudyStruct) SetBuildTime(studyUuid string, buildTime int64) error {
	_, err := s.collection.UpdateOne(s.ctx, bson.M{
		"StudyUuid": studyUuid,
	}, bson.M{
		"$set": bson.M{
			"BuildTime": buildTime,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *StudyStruct) SetUpdatedAt(studyUuid string, updatedAt int64) error {
	_, err := s.collection.UpdateOne(s.ctx, bson.M{
		"StudyUuid": studyUuid,
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

func (s *StudyStruct) SetSync(studyUuid string, status string, syncTime int64) error {
	_, err := s.collection.UpdateOne(s.ctx, bson.M{
		"StudyUuid": studyUuid,
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
