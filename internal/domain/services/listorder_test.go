package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListOrderService_Execute(t *testing.T) {
	// Given
	tmpDir := t.TempDir()

	files := []string{"zebra.txt", "apple.txt", "mango.txt"}
	for _, file := range files {
		_ = os.WriteFile(filepath.Join(tmpDir, file), []byte("content"), 0o600)
	}

	service := NewListOrderService(tmpDir)

	// When
	result := service.Execute()

	// Then
	expectedOrder := []string{"apple.txt", "mango.txt", "zebra.txt"}
	assert.Len(t, result, len(expectedOrder))

	assert.Equal(t, expectedOrder, result)
}
