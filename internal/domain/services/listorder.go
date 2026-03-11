package services

import (
	"os"
	"sort"
)

// ListOrderService - Caso de uso 1: Listar /tmp e ordenar em ordem alfabética.
type ListOrderService struct {
	path string
}

func NewListOrderService(path string) *ListOrderService {
	return &ListOrderService{
		path: path,
	}
}

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
