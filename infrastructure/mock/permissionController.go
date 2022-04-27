package mock

import "github.com/rhodeon/moviescreen/domain/models"

type PermissionController struct {
	Data []userPermission
}

// NewPermissionController creates a PermissionController pointer with the data being
// a copy of the usersPermissions slice to avoid persistent modification across tests.
func NewPermissionController() *PermissionController {
	newUsersPermissions := make([]userPermission, len(usersPermissions))
	copy(newUsersPermissions, usersPermissions)
	return &PermissionController{Data: newUsersPermissions}
}

var permissions = []struct {
	id   int
	code string
}{
	{1, models.PermissionMoviesRead},
	{2, models.PermissionMoviesWrite},
}

type userPermission struct {
	userId       int
	permissionId int
}

var usersPermissions = []userPermission{
	{1, 1},
	{1, 2},
	{2, 1},
	{3, 1},
}

func (p PermissionController) GetAllForUser(user models.User) (models.Permissions, error) {
	permissionIds := []int{}

	for _, userPermission := range usersPermissions {
		if userPermission.userId == user.Id {
			permissionIds = append(permissionIds, userPermission.permissionId)
		}
	}

	perms := models.Permissions{}
	for _, id := range permissionIds {
		for _, permission := range permissions {
			if permission.id == id {
				perms = append(perms, permission.code)
			}
		}
	}

	return perms, nil
}

func (p PermissionController) AddForUser(models.User, ...string) error {
	return nil
}
