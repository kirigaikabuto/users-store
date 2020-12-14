package users_store

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collection *mongo.Collection
)

type userStore struct {
	db *mongo.Database
}

func NewMongoStore(config MongoConfig) (UserStore, error) {
	clientOptions := options.Client().ApplyURI("mongodb://" + config.Host + ":" + config.Port)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.Database)
	collection = db.Collection("users")
	return &userStore{db: db}, nil
}

func (us *userStore) Get(id string) (*User, error) {
	filter := bson.D{{"id", id}}
	user := &User{}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userStore) GetByEmail(email string) (*User, error) {
	filter := bson.D{{"email", email}}
	user := &User{}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userStore) Create(user *User) (*User, error) {
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *userStore) Delete(id string) error {
	filter := bson.D{{"id", id}}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
