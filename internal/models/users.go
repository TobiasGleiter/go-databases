package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ObjectID = primitive.ObjectID

type User struct {
	ID             ObjectID `bson:"_id"`
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// This interface is implemented in the main and enforces that these functions are implemented
type UserModelInterface interface {
	Insert(name, email, password string) error
	Update(name, email, password string) error
	GetByObjID(id ObjectID) (*User, error)
	GetByEmail(email string) (*User, error)
}
