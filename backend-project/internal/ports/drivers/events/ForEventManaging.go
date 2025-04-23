package ports

import entity "exlora/internal/domain/entities"

type ForEventManaging interface {
	GetEventDetails(eventID string) (entity.Event, error)
	GetAllEvents() ([]entity.Event, error)
	CreateEvent(event entity.Event) (entity.Event, error)
	UpdateEvent(event entity.Event) (entity.Event, error)
	DeleteEvent(eventID string) error
	GetEventByID(eventID string) (entity.Event, error)
	GetEventByName(eventName string) (entity.Event, error)
	GetEventByLocation(location string) ([]entity.Event, error)
}
