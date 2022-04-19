package repository

import "github.com/rhodeon/moviescreen/domain/models"

type PermissionRepository interface {
	// GetAllForUser returns the list of permissions granted to the user.
	GetAllForUser(user models.User) (models.Permissions, error)

	// AddForUser grants the specified permission codes to the user.
	AddForUser(user models.User, codes ...string) error
}
