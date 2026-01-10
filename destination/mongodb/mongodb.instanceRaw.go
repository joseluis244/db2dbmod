package mongodb

import (
	"context"

	"github.com/joseluis244/db2dbmod/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func (m *MongoDB) InsertInstanceRawModel(instance models.DestinationInstanceRawType) error {
	collection := m.client.Database(m.db).Collection(models.DestinationInstanceRawCollection)
	opt := options.UpdateOne()
	opt.SetUpsert(true)
	filter := bson.M{"Uuid": instance.Uuid}
	update := bson.M{
		"$set": bson.M{
			"Uuid":      instance.Uuid,
			"Ae":        instance.Ae,
			"Tags":      instance.Tags,
			"StudyUuid": instance.StudyUuid,
			"SerieUuid": instance.SerieUuid,
			"Hash":      instance.Hash,
			"Size":      instance.Size,
			"Path":      instance.Path,
			"Store":     instance.Store,
			"UpdatedAt": instance.UpdatedAt,
		},
		"$setOnInsert": bson.M{
			"CreatedAt": instance.CreatedAt,
			"UpdatedAt": instance.CreatedAt,
		},
	}
	_, err := collection.UpdateOne(context.TODO(), filter, update, opt)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) InsertInstanceRawModels(instances []models.DestinationInstanceRawType) error {
	collection := m.client.Database(m.db).Collection(models.DestinationInstanceRawCollection)
	models := []mongo.WriteModel{}
	for _, instance := range instances {
		model := mongo.NewUpdateOneModel()
		model.SetFilter(bson.M{"Uuid": instance.Uuid})
		model.SetUpdate(bson.M{
			"$set": bson.M{
				"Uuid":      instance.Uuid,
				"Ae":        instance.Ae,
				"Tags":      instance.Tags,
				"StudyUuid": instance.StudyUuid,
				"SerieUuid": instance.SerieUuid,
				"Hash":      instance.Hash,
				"Size":      instance.Size,
				"Path":      instance.Path,
				"Store":     instance.Store,
				"UpdatedAt": instance.UpdatedAt,
			},
			"$setOnInsert": bson.M{
				"CreatedAt": instance.CreatedAt,
				"UpdatedAt": instance.CreatedAt,
			},
		})
		model.SetUpsert(true)
		models = append(models, model)
	}
	_, err := collection.BulkWrite(context.TODO(), models)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoDB) UpdateInstanceRaw(filter bson.M, update bson.M) error {
	collection := m.client.Database(m.db).Collection(models.DestinationInstanceRawCollection)
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
