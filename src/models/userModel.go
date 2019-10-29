package models

import "time"

// User is the structure for accessing table user
type User struct {
	ID           int64
	UUID         string
	Role         *Role
	Account      string
	Password     string
	Name         string
	Email        string
	Status       string
	Score        int
	Verification bool
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	CreatedBy    string
	UpdatedBy    string
}
