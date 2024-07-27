package models

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoUserModel struct {
	DB *mongo.Client
}

func (m *MongoUserModel) Insert(name, email, password string) error {
	fmt.Println("Insert user into mongodb: ", name)
	return nil
}
