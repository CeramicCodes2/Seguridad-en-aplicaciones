package entities

import "time"

// Event represents an event entity in the domain.
type Event struct {
	EventName           string    `json:"event_name"`           // Name of the event
	EventDate           time.Time `json:"event_date"`           // Date and time of the event in ISO 8601 format
	Description         string    `json:"description"`          // Detailed description of the event (optional)
	Location            string    `json:"location"`             // Physical or virtual location (optional)
	OrganizerID         int       `json:"organizer_id"`         // ID of the user who created the event
	CreatedAt           time.Time `json:"created_at"`           // Timestamp of when the event was create
	UsuariosRegistrados []int     `json:"usuarios_registrados"` // List of user IDs registered for the event
}
type EventDashboard struct {
	Nombre            string `json:"nombre"`             // Name of the event
	UsuariosApuntados int    `json:"usuarios_apuntados"` // Number of users attending the event
}
