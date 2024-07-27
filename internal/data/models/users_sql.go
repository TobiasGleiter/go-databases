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
