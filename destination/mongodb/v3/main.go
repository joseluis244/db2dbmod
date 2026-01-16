package v3

import (
	"context"

	"github.com/joseluis244/db2dbmod/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type V3Struct struct {
	client     *mongo.Client
	db         string
	collection *mongo.Collection
	ctx        context.Context
}

func New(client *mongo.Client, db string, collection string) *V3Struct {
	coll := client.Database(db).Collection(collection)
	return &V3Struct{
		client:     client,
		db:         db,
		collection: coll,
		ctx:        context.TODO(),
	}
}

func (v *V3Struct) UpsertV3s(V3s []models.DestinationV3Type) error {
	Models := []mongo.WriteModel{}
	for _, V3 := range V3s {
		Model := mongo.NewUpdateOneModel().SetFilter(bson.M{"StudyUuid": V3.StudyUuid}).SetUpdate(bson.M{"$set": V3}).SetUpsert(true)
		Models = append(Models, Model)
		if len(Models) == 500 {
			_, err := v.collection.BulkWrite(v.ctx, Models)
			if err != nil {
				return err
			}
			Models = []mongo.WriteModel{}
		}
	}
	if len(Models) > 0 {
		_, err := v.collection.BulkWrite(v.ctx, Models)
		if err != nil {
			return err
		}
	}
	return nil
}
