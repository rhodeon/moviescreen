// Package repository provides generic interfaces to be implemented for different sources
// (like mocks and the database).
package repository

// Repositories encapsulates all available repositories for easy reuse.
type Repositories struct {
	Movies MovieRepository
}
