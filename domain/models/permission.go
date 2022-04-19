package models

type Permissions []string

const (
	PermissionMoviesRead  = "movies:read"
	PermissionMoviesWrite = "movies:write"
	PermissionMetricsView = "metrics:view"
)

// Includes returns true if the specified code is amongst the permissions,
// and false otherwise
func (p Permissions) Includes(code string) bool {
	for _, permissions := range p {
		if permissions == code {
			return true
		}
	}
	return false
}
