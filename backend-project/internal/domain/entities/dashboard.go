package entities

import "time"

// Dashboard represents the dashboard data structure.
type Dashboard struct {
	TotalUsuarios     int                 `json:"total_usuarios"`     // Total number of users
	TotalEventos      int                 `json:"total_eventos"`      // Total number of events
	ActividadReciente []ActividadReciente `json:"actividad_reciente"` // Recent activity data
	Eventos           []EventDashboard    `json:"eventos"`            // List of events with user counts
}

// ActividadReciente represents recent activity on the dashboard.
type ActividadReciente struct {
	Fecha     time.Time `json:"fecha"`     // Date of the activity
	Eventos   int       `json:"eventos"`   // Number of events on that date
	Registros int       `json:"registros"` // Number of user registrations on that date
}
