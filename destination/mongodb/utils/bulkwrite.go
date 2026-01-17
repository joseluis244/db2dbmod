package utils

import (
	"context"

	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func BulkWrite(ctx context.Context, collection *mongo.Collection, models []mongo.WriteModel) (*mongo.BulkWriteResult, error) {
	if len(models) == 0 {
		return &mongo.BulkWriteResult{}, nil
	}

	const batchSize = 500
	opts := options.BulkWrite().SetOrdered(false)

	var totalResult mongo.BulkWriteResult

	for i := 0; i < len(models); i += batchSize {
		end := i + batchSize
		if end > len(models) {
			end = len(models)
		}

		result, err := collection.BulkWrite(ctx, models[i:end], opts)
		if err != nil {
			return &totalResult, fmt.Errorf("bulk write failed at batch %d-%d: %w", i, end, err)
		}

		// Acumular resultados
		if result != nil {
			totalResult.InsertedCount += result.InsertedCount
			totalResult.MatchedCount += result.MatchedCount
			totalResult.ModifiedCount += result.ModifiedCount
			totalResult.DeletedCount += result.DeletedCount
			totalResult.UpsertedCount += result.UpsertedCount
		}
	}

	return &totalResult, nil
}
