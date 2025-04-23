package entities

// Role represents a role in the application.
type Role struct {
	ID          int          `json:"id"`          // Unique identifier for the role
	Name        string       `json:"name"`        // Name of the role (e.g., "admin", "user")
	Description string       `json:"description"` // Description of the role
	Permissions []Permission `json:"permissions"` // List of permissions associated with the role
}
