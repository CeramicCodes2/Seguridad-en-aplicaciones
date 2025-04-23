package entities

import "time"

// User represents a user entity in the domain.
type User struct {
	Username  string    `json:"username"`   // Username of the user
	Email     string    `json:"email"`      // Email address of the user
	Password  string    `json:"password"`   // Hashed password
	CreatedAt time.Time `json:"created_at"` // Timestamp of when the user was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp of the last update
	Role      int       `json:"role"`       // Role of the user (e.g., admin, user, etc.)
}

// UserRepository represents the persisted user entity.
type UserRepository struct {
	ID        string    `json:"id"`         // Unique identifier for the user
	Email     string    `json:"email"`      // Email address of the user
	Password  string    `json:"password"`   // Hashed password
	Username  string    `json:"username"`   // Optional username
	CreatedAt time.Time `json:"created_at"` // Timestamp of when the user was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp of the last update
}
