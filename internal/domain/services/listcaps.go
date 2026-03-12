package services

import (
	"os"
	"strings"
)

// ListCapsService implements the use case to list files and convert names to uppercase.
type ListCapsService struct {
	path string
}

// NewListCapsService creates a new ListCapsService instance.
func NewListCapsService(path string) *ListCapsService {
	return &ListCapsService{
		path: path,
	}
}

// Execute lists files in the specified directory and returns them in uppercase.
func (s *ListCapsService) Execute() []string {
	files, err := os.ReadDir(s.path)
	if err != nil {
		panic(err)
	}

	result := make([]string, 0, len(files))
	for _, file := range files {
		result = append(result, strings.ToUpper(file.Name()))
	}

	return result
}
