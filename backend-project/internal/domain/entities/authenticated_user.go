package entities

// AuthenticatedUser represents a user with authentication details.
type AuthenticatedUser struct {
	User        User         `json:"user"`        // Embedded User entity
	Roles       []Role       `json:"roles"`       // List of roles associated with the user
	Permissions []Permission `json:"permissions"` // List of permissions associated with the user
	Token       string       `json:"token"`       // Authentication token (e.g., JWT)
}
