package repository

import "github.com/rhodeon/moviescreen/domain/models"

type PermissionRepository interface {
	// GetAllForUser returns the list of permissions granted to a user.
	GetAllForUser(user models.User) (models.Permissions, error)
}
