package models

import (
	"database/sql"
	"fmt"
)

type ID = int

type SQLUserModel struct {
	DB *sql.DB
}

func (m *SQLUserModel) Insert(name, email string, password string) error {
	fmt.Println("Insert user into sql database: ", name)
	return nil
}

func (m *SQLUserModel) UpdateByObjId(id ObjectID, name, email, password string) error {
	fmt.Println("Update user in sql database: ", name)
	return nil
}

func (m *SQLUserModel) GetByEmail(email string) (*User, error) {
	var user User
	return &user, nil
}

// Refactor: ObjectID isn't used by sql database, it's a mongodb thing.
func (m *SQLUserModel) GetByObjID(id ObjectID) (*User, error) {
	var user User
	return &user, nil
}
