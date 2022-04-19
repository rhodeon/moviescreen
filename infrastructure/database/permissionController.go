package database

import (
	"context"
	"database/sql"
	"github.com/rhodeon/moviescreen/domain/models"
	"time"
)

type PermissionController struct {
	Db *sql.DB
}

func (p PermissionController) GetAllForUser(user models.User) (models.Permissions, error) {
	// join the permissions, users_permissions and users table to retrieve the
	// permission codes for the user
	stmt := `SELECT permissions.code 
	FROM permissions
	INNER JOIN users_permissions ON permissions.id = users_permissions.permission_id
	INNER JOIN users ON users_permissions.user_id = users.id
	WHERE users.id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := p.Db.QueryContext(ctx, stmt, user.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	permissions := models.Permissions{}
	for rows.Next() {
		var permission string
		_ = rows.Scan(&permission)
		permissions = append(permissions, permission)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return permissions, nil
}
