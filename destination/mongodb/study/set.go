package study

import (
	"github.com/joseluis244/db2dbmod/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (s *StudyStruct) UpsertStudy(study models.DestinationStudyType) error {
	opt := options.UpdateOne()
	opt.SetUpsert(true)

	_, err := s.collection.UpdateOne(s.ctx, bson.M{
		"StudyUuid": study.StudyUuid},
		bson.M{
			"$setOnInsert": bson.M{
				"DealerID":  study.DealerID,
				"ClientID":  study.ClientID,
				"BranchID":  study.BranchID,
				"StudyUuid": study.StudyUuid,
				"CreatedAt": study.CreatedAt,
				"BuildTime": 0,
				"Sync": models.SyncType{
					Status:   "pending",
					SyncTime: 0,
				},
			},
			"$set": bson.M{
				"Tags":      study.Tags,
				"UpdatedAt": study.UpdatedAt,
			},
		},
		opt)
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
