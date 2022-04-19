package mock

import "github.com/rhodeon/moviescreen/domain/models"

var permissions = []struct {
	id   int
	code string
}{
	{1, models.PermissionMoviesRead},
	{2, models.PermissionMoviesWrite},
}

var usersPermissions = []struct {
	userId       int
	permissionId int
}{
	{1, 1},
	{1, 2},
	{2, 1},
	{3, 1},
}

type PermissionController struct{}

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

func (p PermissionController) AddForUser(user models.User, codes ...string) error {
	return nil
}
