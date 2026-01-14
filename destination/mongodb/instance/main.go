package instance

import (
	"context"

	"github.com/joseluis244/db2dbmod/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type InstanceStruct struct {
	client     *mongo.Client
	db         string
	collection *mongo.Collection
	ctx        context.Context
}

func New(client *mongo.Client, db string, collection string) *InstanceStruct {
	return &InstanceStruct{
		client:     client,
		db:         db,
		collection: client.Database(db).Collection(collection),
		ctx:        context.TODO(),
	}
}

func (i *InstanceStruct) UpsertInstance(instance models.DestinationInstanceType) error {
	opt := options.UpdateOne()
	opt.SetUpsert(true)

	_, err := i.collection.UpdateOne(i.ctx, bson.M{
		"StudyUuid": instance.StudyUuid,
		"SerieUuid": instance.SerieUuid,
		"Uuid":      instance.Uuid,
	}, bson.M{
		"$setOnInsert": bson.M{
			"DealerID":  instance.DealerID,
			"ClientID":  instance.ClientID,
			"BranchID":  instance.BranchID,
			"StudyUuid": instance.StudyUuid,
			"SerieUuid": instance.SerieUuid,
			"Uuid":      instance.Uuid,
			"Ae":        instance.Ae,
			"Tags":      instance.Tags,
			"CreatedAt": instance.CreatedAt,
			"BuildTime": 0,
			"Sync": models.SyncType{
				Status:   "pending",
				SyncTime: 0,
			},
		},
		"$set": bson.M{
			"Tags":      instance.Tags,
			"UpdatedAt": instance.UpdatedAt,
		},
	},
		opt)
	if err != nil {
		return err
	}
	return nil
}

func (i *InstanceStruct) SetUpdatedAt(uuid string, updatedAt int64) error {
	_, err := i.collection.UpdateOne(i.ctx, bson.M{
		"Uuid": uuid,
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

func (i *InstanceStruct) SetSync(uuid string, status string, syncTime int64) error {
	_, err := i.collection.UpdateOne(i.ctx, bson.M{
		"Uuid": uuid,
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
