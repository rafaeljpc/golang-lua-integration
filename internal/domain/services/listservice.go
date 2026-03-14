// Package services provides domain business logic for use cases.
package services

// ListService defines the interface for listing services.
type ListService interface {
	Execute() []string
}
