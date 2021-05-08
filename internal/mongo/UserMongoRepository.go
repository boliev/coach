package mongo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/boliev/coach/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserMongoRepository(client *mongo.Client, db string, coll string) *UserMongoRepository {
	collection := client.Database(db).Collection(coll)
	return &UserMongoRepository{
		client:     client,
		collection: collection,
	}
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := r.collection.InsertOne(ctx, user)
	return res, err
}

func (r UserMongoRepository) FindAll() ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println(err.Error())
	}
	defer cursor.Close(ctx)

	return r.cursorToDomain(cursor, ctx), nil
}

func (r UserMongoRepository) Find(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("there is no such user")
	}

	var result bson.D
	r.collection.FindOne(ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&result)

	return r.bsonToDomain(result.Map()), nil
}

func (r UserMongoRepository) Delete(id string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}

	r.collection.DeleteOne(ctx, bson.D{{Key: "_id", Value: objectId}})
}
