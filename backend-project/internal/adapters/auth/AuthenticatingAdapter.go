package adapters

import (
	entity "explora/internal/domain/entities"
)

type AuthenticatingAdapter struct {
}

func Authenticate(username, password string) (entity.AuthenticatedUser, error) {
	return entity.AuthenticatedUser{}, nil

}
