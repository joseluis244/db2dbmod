package system

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type SystemStruct struct {
	client     *mongo.Client
	db         string
	collection *mongo.Collection
	ctx        context.Context
}

func New(client *mongo.Client, db string, collection string) *SystemStruct {
	coll := client.Database(db).Collection(collection)
	return &SystemStruct{
		client:     client,
		db:         db,
		collection: coll,
		ctx:        context.TODO(),
	}
}

func (s *SystemStruct) SetLastChange(lastChange int64) error {
	filter := bson.M{"id": 0}
	update := bson.M{"$set": bson.M{"lastchange": lastChange}}
	opt := options.UpdateOne()
	opt.SetUpsert(true)
	_, err := s.collection.UpdateOne(s.ctx, filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}

func (s *SystemStruct) GetLastChange() (int64, error) {
	var result bson.M
	err := s.collection.FindOne(s.ctx, bson.M{"id": 0}).Decode(&result)
	if err != nil {
		return 0, err
	}
	return result["lastchange"].(int64), nil
}
