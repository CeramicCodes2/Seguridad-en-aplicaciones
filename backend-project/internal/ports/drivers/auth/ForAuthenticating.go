package ports

import entity "exlora/internal/domain/entities"

type ForAuthenticating interface {
	Authenticate(username, password string) (entity.AuthenticatedUser, error)
	// Authenticate authenticates a user with the given username and password.
}
