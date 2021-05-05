package repository

import (
	"context"
	"time"

	"github.com/boliev/coach/internal/domain"
	"github.com/boliev/coach/pkg/mongo_factory"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserMongoRepository() *UserMongoRepository {
	client := mongo_factory.NewClient()
	collection := client.Database("coach").Collection("user")
	return &UserMongoRepository{
		client:     client,
		collection: collection,
	}
}

func (r UserMongoRepository) Create(user *domain.User) (interface{}, error) {
	ctx, cancel := r.context()
	defer cancel()

	res, err := r.collection.InsertOne(ctx, user)
	return res, err
}

func (u UserMongoRepository) context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
