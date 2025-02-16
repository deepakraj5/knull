package entities

import "time"

type User struct {
	ID        uint
	Name      string
	Email     string
	IsActive  bool
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
