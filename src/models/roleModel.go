package models

import "time"

// Role is the structure for accessing table role
type Role struct {
	ID        int64
	UUID      string
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt *time.Time
	CreatedBy string
	UpdatedBy string
}
