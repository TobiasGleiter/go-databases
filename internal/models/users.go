package models

import (
	"database/sql"
	"fmt"
	"time"
)

type UserModelInterface interface {
	Insert(name, email, password string) error
}

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email string, password string) error {
	fmt.Println("Insert user into database: ", name)
	return nil
}
