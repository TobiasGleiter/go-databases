package models

import (
	"database/sql"
	"fmt"
)

type SQLUserModel struct {
	DB *sql.DB
}

func (m *SQLUserModel) Insert(name, email string, password string) error {
	fmt.Println("Insert user into sql database: ", name)
	return nil
}

func (m *SQLUserModel) Update(name, email, password string) error {
	fmt.Println("Update user in sql database: ", name)
	return nil
}
