// Package services provides domain business logic for use cases.
package services

import (
	"os"
	"sort"
)

// ListOrderService implements the use case to list and order files alphabetically.
type ListOrderService struct {
	path string
}

var _ ListService = (*ListOrderService)(nil)

// NewListOrderService creates a new ListOrderService instance.
func NewListOrderService(path string) *ListOrderService {
	return &ListOrderService{
		path: path,
	}
}

// Execute lists files in the specified directory and returns them in alphabetical order.
func (s *ListOrderService) Execute() []string {
	files, err := os.ReadDir(s.path)
	if err != nil {
		panic(err)
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	result := make([]string, 0, len(files))
	for _, file := range files {
		result = append(result, file.Name())
	}

	return result
}
