package services

import (
	"os"
	"strings"
)

// ListCapsService - Caso de uso 2: Listar /tmp e colocar tudo em maiusculo e salvar resultado no /tmp.
type ListCapsService struct {
	path string
}

func NewListCapsService(path string) *ListCapsService {
	return &ListCapsService{
		path: path,
	}
}

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
