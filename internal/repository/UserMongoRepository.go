package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/boliev/coach/internal/domain"
	"github.com/boliev/coach/pkg/mongo_factory"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (u UserMongoRepository) context() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

func (u UserMongoRepository) bsonToDomain(mapUser map[string]interface{}) *domain.User {
	return &domain.User{
		Email:    mapUser["email"].(string),
		Password: mapUser["password"].(string),
		Id:       mapUser["_id"].(primitive.ObjectID).Hex(),
	}
}

func (u UserMongoRepository) cursorToDomain(cursor *mongo.Cursor, ctx context.Context) []*domain.User {
	var users []*domain.User
	for cursor.Next(ctx) {
		var result bson.D
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, u.bsonToDomain(result.Map()))
	}

	return users
}

func (r UserMongoRepository) Create(user *domain.User) (interface{}, error) {
	ctx, cancel := r.context()
	defer cancel()

	res, err := r.collection.InsertOne(ctx, user)
	return res, err
}

func (r UserMongoRepository) FindAll() ([]*domain.User, error) {
	ctx, cancel := r.context()
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err.Error())
	}
	defer cursor.Close(ctx)

	return r.cursorToDomain(cursor, ctx), nil
}
