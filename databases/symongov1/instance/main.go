package instance

import (
	"context"

	"github.com/joseluis244/db2dbmod/databases/symongov1/models"
	"github.com/joseluis244/db2dbmod/databases/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func createInstance(InstanceUuid string, Ae string, SerieUuid string, StudyUuid string, updatedAt int64, Tags map[string]interface{}, hash string, size int64, path string) bson.M {
	return bson.M{
		"$setOnInsert": bson.M{
			"Uuid":      InstanceUuid,
			"Ae":        Ae,
			"Id":        4,
			"SerieUuid": SerieUuid,
			"StudyUuid": StudyUuid,
		},
		"$set": bson.M{
			"CloudSync": 0,
			"Hash":      hash,
			"Path":      path,
			"Size":      size,
			"Update":    updatedAt,
			"tags":      Tags,
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

func (i *InstanceStruct) UpsertInstances(instances []models.SyMongoV1InstanceType) error {
	Models := []mongo.WriteModel{}
	for _, instance := range instances {
		filter := bson.M{"Uuid": instance.Uuid}
		update := createInstance(instance.Uuid, instance.Ae, instance.SerieUuid, instance.StudyUuid, instance.Update, instance.Tags, instance.Hash, instance.Size, instance.Path)
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

func (i *InstanceStruct) UpsertInstance(instance models.SyMongoV1InstanceType) error {
	opt := options.UpdateOne()
	opt.SetUpsert(true)
	filter := bson.M{"Uuid": instance.Uuid}
	update := createInstance(instance.Uuid, instance.Ae, instance.SerieUuid, instance.StudyUuid, instance.Update, instance.Tags, instance.Hash, instance.Size, instance.Path)
	_, err := i.collection.UpdateOne(i.ctx, filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}

func (i *InstanceStruct) SetUpdatedAt(uuid string, updatedAt int64) error {
	filter := bson.M{"Uuid": uuid}
	update := bson.M{"$set": bson.M{"Update": updatedAt}}
	_, err := i.collection.UpdateOne(i.ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
