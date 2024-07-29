package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	testDatabase    = "test"
	usersCollection = "users"
)

type MongoUserModel struct {
	DB *mongo.Client
}

func (m *MongoUserModel) Insert(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	user := struct {
		Name           string
		Email          string
		HashedPassword []byte
		Created        time.Time
	}{
		Name:           name,
		Email:          email,
		HashedPassword: hashedPassword,
		Created:        time.Now(),
	}

	usersCollection := m.DB.Database(testDatabase).Collection(usersCollection)
	_, err = usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoUserModel) Update(objId ObjectID, name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: name},
		{Key: "email", Value: email},
		{Key: "hashedpassword", Value: hashedPassword},
	}}}

	usersCollection := m.DB.Database(testDatabase).Collection(usersCollection)
	filter := bson.D{{Key: "_id", Value: objId}}
	_, err = usersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m *MongoUserModel) GetByObjID(objId ObjectID) (*User, error) {
	usersCollection := m.DB.Database("test").Collection("users")
	filter := bson.D{{Key: "_id", Value: objId}}

	var user User
	err := usersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

func (m *MongoUserModel) GetByEmail(email string) (*User, error) {
	usersCollection := m.DB.Database("test").Collection("users")
	filter := bson.D{{Key: "email", Value: email}}

	var user User
	err := usersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	return &user, nil
}

func (m *MongoUserModel) DeleteByObjId(objId ObjectID) error {
	usersCollection := m.DB.Database(testDatabase).Collection(usersCollection)
	filter := bson.D{{Key: "_id", Value: objId}}

	_, err := usersCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}
