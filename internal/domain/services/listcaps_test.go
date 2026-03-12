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
	// Given
	tmpDir := t.TempDir()

	files := []string{"zebra.txt", "apple.txt", "mango.txt"}
	for _, file := range files {
		err := os.WriteFile(filepath.Join(tmpDir, file), []byte("content"), 0o600)
		require.NoError(t, err)
	}

	service := NewListCapsService(tmpDir)

	// When
	result := service.Execute()

	// Then - Verify files are converted to uppercase
	expectedFiles := []string{"ZEBRA.TXT", "APPLE.TXT", "MANGO.TXT"}
	assert.Len(t, result, len(expectedFiles))

	for _, expected := range expectedFiles {
		assert.Contains(t, result, expected)
	}
}

func TestListCapsService_Execute_EmptyDirectory(t *testing.T) {
	t.Parallel()
	// Given
	tmpDir := t.TempDir()
	service := NewListCapsService(tmpDir)

	// When
	result := service.Execute()

	// Then - Verify empty directory returns empty result
	assert.Empty(t, result)
	assert.Empty(t, result)
}

func TestListCapsService_Execute_MixedCaseFilenames(t *testing.T) {
	t.Parallel()
	// Given
	tmpDir := t.TempDir()

	files := []string{"CamelCase.txt", "lowercase.txt", "UPPERCASE.TXT", "MixedCase.Txt"}
	for _, file := range files {
		err := os.WriteFile(filepath.Join(tmpDir, file), []byte("content"), 0o600)
		require.NoError(t, err)
	}

	service := NewListCapsService(tmpDir)

	// When
	result := service.Execute()

	// Then - Verify all filenames are converted to uppercase
	expectedFiles := []string{"CAMELCASE.TXT", "LOWERCASE.TXT", "UPPERCASE.TXT", "MIXEDCASE.TXT"}
	assert.Len(t, result, len(expectedFiles))

	for _, expected := range expectedFiles {
		assert.Contains(t, result, expected)
	}
}

func TestListCapsService_Execute_WithSubdirectories(t *testing.T) {
	t.Parallel()
	// Given
	tmpDir := t.TempDir()

	// Create files in root
	err := os.WriteFile(filepath.Join(tmpDir, "rootfile.txt"), []byte("content"), 0o600)
	require.NoError(t, err)

	// Create subdirectory
	subDir := filepath.Join(tmpDir, "subdir")
	err = os.Mkdir(subDir, 0o750)
	require.NoError(t, err)

	service := NewListCapsService(tmpDir)

	// When
	result := service.Execute()

	// Then - Verify root file and directory are converted to uppercase
	assert.Contains(t, result, "ROOTFILE.TXT")
	assert.Contains(t, result, "SUBDIR")
}

func TestListCapsService_Execute_WithSpecialCharacters(t *testing.T) {
	t.Parallel()
	// Given
	tmpDir := t.TempDir()

	files := []string{"file-with-dash.txt", "file_with_underscore.txt", "file.multiple.dots.txt"}
	for _, file := range files {
		err := os.WriteFile(filepath.Join(tmpDir, file), []byte("content"), 0o600)
		require.NoError(t, err)
	}

	service := NewListCapsService(tmpDir)

	// When
	result := service.Execute()

	// Then - Verify special characters are preserved
	expectedFiles := []string{"FILE-WITH-DASH.TXT", "FILE_WITH_UNDERSCORE.TXT", "FILE.MULTIPLE.DOTS.TXT"}
	assert.Len(t, result, len(expectedFiles))

	for _, expected := range expectedFiles {
		assert.Contains(t, result, expected)
	}
}

func TestListCapsService_Execute_InvalidDirectory(t *testing.T) {
	t.Parallel()
	// Given
	service := NewListCapsService("/nonexistent/directory/path")

	// When & Then - Verify panic on invalid directory
	assert.Panics(t, func() {
		service.Execute()
	})
}
