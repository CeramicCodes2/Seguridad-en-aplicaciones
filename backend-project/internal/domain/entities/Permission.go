package entities

// Permission represents a permission in the application.
type Permission struct {
	ID          int    `json:"id"`          // Unique identifier for the permission
	Name        string `json:"name"`        // Name of the permission
	Description string `json:"description"` // Description of what the permission allows
}
