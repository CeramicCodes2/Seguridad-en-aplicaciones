package ports

import (
	entity "exlora/internal/domain/entities"
)

type ForAuthDetails interface {
	getAuthDetails() (entity.AuthenticatedUser, error)
	getUserPermissions(userID string) ([]entity.Permission, error)
	getUserRoles(userID string) ([]entity.Role, error)
}
