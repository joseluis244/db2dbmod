package v3

import (
	"context"

	"github.com/joseluis244/db2dbmod/databases/symongov2/models"
	"github.com/joseluis244/db2dbmod/databases/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type V3Struct struct {
	client     *mongo.Client
	db         string
	collection *mongo.Collection
	ctx        context.Context
}

func New(client *mongo.Client, db string, collection string) *V3Struct {
	coll := client.Database(db).Collection(collection)

	coll.Indexes().CreateMany(context.TODO(), []mongo.IndexModel{
		{
			Keys:    bson.M{"StudyUuid": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.M{"Tags.0008,0020": 1},
			Options: options.Index().SetUnique(false),
		},
	})
	return &V3Struct{
		client:     client,
		db:         db,
		collection: coll,
		ctx:        context.TODO(),
	}
}

func (v *V3Struct) UpsertV3s(V3s []models.SyMongoV2V3Type) error {
	Models := []mongo.WriteModel{}
	for _, V3 := range V3s {
		filter := bson.M{"StudyUuid": V3.StudyUuid}
		update := bson.M{"$set": V3}
		Model := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
		Models = append(Models, Model)
	}
	_, err := utils.BulkWrite(v.ctx, v.collection, Models)
	if err != nil {
		return err
	}
	return nil
}
