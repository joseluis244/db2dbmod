package instance

import (
	"context"

	"github.com/joseluis244/db2dbmod/databases/symongov2/models"
	"github.com/joseluis244/db2dbmod/databases/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func createInstance(DealerID string, ClientID string, BranchID string, StudyUuid string, SerieUuid string, InstanceUuid string, Ae string, CreatedAt int64, Tags map[string]interface{}, updatedAt int64, hash string, size int64, path string, store string) bson.M {
	return bson.M{
		"$setOnInsert": bson.M{
			"DealerID":     DealerID,
			"ClientID":     ClientID,
			"BranchID":     BranchID,
			"StudyUuid":    StudyUuid,
			"SerieUuid":    SerieUuid,
			"InstanceUuid": InstanceUuid,
			"Ae":           Ae,
			"CreatedAt":    CreatedAt,
			"Sync": models.SyncType{
				Status:   "pending",
				SyncTime: 0,
			},
		},
		"$set": bson.M{
			"Tags":      Tags,
			"UpdatedAt": updatedAt,
			"Hash":      hash,
			"Size":      size,
			"Path":      path,
			"Store":     store,
		},
	}
}

type InstanceStruct struct {
	client     *mongo.Client
	db         string
	collection *mongo.Collection
	ctx        context.Context
}

func New(client *mongo.Client, db string, collection string) *InstanceStruct {
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
		{
			Keys:    bson.M{"InstanceUuid": 1},
			Options: options.Index().SetUnique(false),
		},
	})
	return &InstanceStruct{
		client:     client,
		db:         db,
		collection: coll,
		ctx:        context.TODO(),
	}
}

func (i *InstanceStruct) UpsertInstances(instances []models.DestinationInstanceType) error {
	Models := []mongo.WriteModel{}
	for _, instance := range instances {
		filter := bson.M{"InstanceUuid": instance.InstanceUuid}
		update := createInstance(instance.DealerID, instance.ClientID, instance.BranchID, instance.StudyUuid, instance.SerieUuid, instance.InstanceUuid, instance.Ae, instance.CreatedAt, instance.Tags, instance.UpdatedAt, instance.Hash, instance.Size, instance.Path, instance.Store)
		Model := mongo.NewUpdateOneModel()
		Model.SetFilter(filter)
		Model.SetUpdate(update)
		Model.SetUpsert(true)
		Models = append(Models, Model)
	}
	_, err := utils.BulkWrite(i.ctx, i.collection, Models)
	if err != nil {
		return err
	}
	return nil
}

func (i *InstanceStruct) UpsertInstance(instance models.DestinationInstanceType) error {
	opt := options.UpdateOne()
	opt.SetUpsert(true)
	filter := bson.M{"InstanceUuid": instance.InstanceUuid}
	update := createInstance(instance.DealerID, instance.ClientID, instance.BranchID, instance.StudyUuid, instance.SerieUuid, instance.InstanceUuid, instance.Ae, instance.CreatedAt, instance.Tags, instance.UpdatedAt, instance.Hash, instance.Size, instance.Path, instance.Store)
	_, err := i.collection.UpdateOne(i.ctx, filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}

func (i *InstanceStruct) SetUpdatedAt(uuid string, updatedAt int64) error {
	filter := bson.M{"InstanceUuid": uuid}
	update := bson.M{"$set": bson.M{"UpdatedAt": updatedAt}}
	_, err := i.collection.UpdateOne(i.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (i *InstanceStruct) SetSync(uuid string, status string, syncTime int64) error {
	filter := bson.M{"InstanceUuid": uuid}
	update := bson.M{"$set": bson.M{"Sync": models.SyncType{
		Status:   status,
		SyncTime: syncTime,
	}}}
	_, err := i.collection.UpdateOne(i.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
