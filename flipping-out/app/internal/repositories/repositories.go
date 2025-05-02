package repositories

import (
	"context"

	"github.com/joshbrgs/flipping-out/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type FeatureFlagRepository interface {
	GetAll(ctx context.Context) ([]models.FeatureFlag, error)
	GetByID(ctx context.Context, id string) (*models.FeatureFlag, error)
	Create(ctx context.Context, flag models.FeatureFlag) error
	Update(ctx context.Context, id string, update bson.M) error
	Delete(ctx context.Context, id string) error
}

type featureFlagRepository struct {
	collection *mongo.Collection
}

func NewFeatureFlagRepository(client *mongo.Client, dbName, collectionName string) FeatureFlagRepository {
	return &featureFlagRepository{
		collection: client.Database(dbName).Collection(collectionName),
	}
}

func (r *featureFlagRepository) GetAll(ctx context.Context) ([]models.FeatureFlag, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var flags []models.FeatureFlag
	if err := cursor.All(ctx, &flags); err != nil {
		return nil, err
	}
	return flags, nil
}

func (r *featureFlagRepository) GetByID(ctx context.Context, id string) (*models.FeatureFlag, error) {
	var flag models.FeatureFlag
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&flag)
	return &flag, err
}

func (r *featureFlagRepository) Create(ctx context.Context, flag models.FeatureFlag) error {
	_, err := r.collection.InsertOne(ctx, flag)
	return err
}

func (r *featureFlagRepository) Update(ctx context.Context, id string, update bson.M) error {
	_, err := r.collection.UpdateByID(ctx, id, bson.M{"$set": update})
	return err
}

func (r *featureFlagRepository) Delete(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
