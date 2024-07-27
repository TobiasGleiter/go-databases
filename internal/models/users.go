package models

import (
	"time"
)

type User struct {
	ID             int
	Name           string
	Email          string
	HashedPassword []byte
	Created        time.Time
}

// This interface is implemented in the main and enforces that these functions are implemented
type UserModelInterface interface {
	Insert(name, email, password string) error
	Update(name, email, password string) error
}
