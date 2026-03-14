package commands

import (
	"bytes"
	"context"
	"golang-lua-integration/internal/domain/services"
	"io"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// TestListCapsCommandExecuteSuccess tests successful execution of ListCapsCommand with mocked service.
func TestListCapsCommandExecuteSuccess(t *testing.T) {
	// Given
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockService := services.NewMockListService(ctrl)
	expectedListFileName := []string{"FILE1.TXT", "FILE2.TXT"}

	mockService.EXPECT().Execute().Return(expectedListFileName)

	cmd := NewListCapsCommand(ctx, mockService)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	// When
	err := cmd.Run(ctx, &cmd.Command)

	// Then
	var buf bytes.Buffer
	go func() {
		defer wg.Done()
		io.Copy(&buf, r)
	}()
	w.Close()
	os.Stdout = oldStdout

	wg.Wait()

	resultStr := buf.String()

	assert.NoError(t, err)
	assert.Equal(t, strings.Join(expectedListFileName, "\n")+"\n", resultStr)
}

// TestListCapsCommandNameAndDescription tests the command metadata.
func TestListCapsCommandNameAndDescription(t *testing.T) {
	// Given
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockService := services.NewMockListService(ctrl)

	// When
	cmd := NewListCapsCommand(ctx, mockService)

	// Then
	assert.Equal(t, "list-caps", cmd.Name)
	assert.Contains(t, cmd.Description, "uppercase")
}

// TestListCapsCommandContextCancellation tests context cancellation handling.
func TestListCapsCommandContextCancellation(t *testing.T) {
	// Given
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockService := services.NewMockListService(ctrl)

	// When
	cmd := NewListCapsCommand(ctx, mockService)

	_, cancel := context.WithCancel(ctx)
	cancel()

	// Then
	require.NotNil(t, cmd)
	require.Equal(t, "list-caps", cmd.Name)
	assert.NotNil(t, cmd.Action)
}

// TestListCapsCommandActionAssignment tests that Action is properly set.
func TestListCapsCommandActionAssignment(t *testing.T) {
	// Given
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockService := services.NewMockListService(ctrl)

	// When
	cmd := NewListCapsCommand(ctx, mockService)

	// Then
	require.NotNil(t, cmd)
	require.NotNil(t, cmd.Action)
	assert.Equal(t, "list-caps", cmd.Name)
}

// TestListCapsCommandNoServiceInjection tests error handling when service is not injected.
func TestListCapsCommandNoServiceInjection(t *testing.T) {
	// Given
	ctx := t.Context()
	cmd := NewListCapsCommand(ctx, nil)

	// When
	err := cmd.Run(ctx, nil)

	// Then
	require.Error(t, err)
}
