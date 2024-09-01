package main

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Email          string
	HashedPassword []byte
	DateOfBirth    string
	BirthPlace     string
	Friends        []uuid.UUID
}
