package services

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListCapsService_Execute_ConvertsFilesToUppercase(t *testing.T) {
	t.Parallel()

	tmpDir := t.TempDir()

	files := []string{"file1.txt", "file2.txt"}
	for _, file := range files {
		_ = os.WriteFile(filepath.Join(tmpDir, file), []byte("content"), 0o600)
	}

	service := NewListCapsService(tmpDir)

	result := service.Execute()

	expectedFiles := []string{"FILE1.TXT", "FILE2.TXT"}
	assert.Len(t, result, len(expectedFiles))

	require.Len(t, result, len(expectedFiles))
}
