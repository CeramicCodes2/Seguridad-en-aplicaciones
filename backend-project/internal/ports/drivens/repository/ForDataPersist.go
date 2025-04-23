package ports

import entity "exlora/internal/domain/entities"

type ForDataPersist interface {
	InsertUser(user entity.User) error
	UpdateUser(user entity.User) error
	DeleteUser(userID string) error
	GetUserByID(userID string) (entity.User, error)
	GetUserByUsername(username string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
}
