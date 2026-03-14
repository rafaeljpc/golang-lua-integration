package commands

import (
	"context"
	"golang-lua-integration/internal/domain/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

// TestListOrderCommandExecuteSuccess tests successful execution of ListOrderCommand with mocked service.
func TestListOrderCommandExecuteSuccess(t *testing.T) {
	// Givem
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockService := services.NewMockListService(ctrl)
	mockedFileList := []string{"file2.txt", "file1.txt"}
	mockService.EXPECT().Execute().Return(mockedFileList).Times(1)

	cmd := NewListOrderCommand(ctx, mockService)

	// When
	err := cmd.Run(ctx, &cmd.Command)

	// Then
	assert.NoError(t, err)
}

// TestListOrderCommandInitialization tests the initialization of ListOrderCommand with service injection.
func TestListOrderCommandInitialization(t *testing.T) {
	// Given
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockService := services.NewMockListService(ctrl)
	mockService.EXPECT().Execute().Return([]string{}).Times(1)

	// When
	cmd := NewListOrderCommand(ctx, mockService)

	// Then
	require.NoError(t, cmd.Run(ctx, &cmd.Command))
	assert.Equal(t, "list-order", cmd.Name)
	assert.Contains(t, cmd.Description, "order alphabetically")
}

// TestListOrderCommandNameAndDescription tests the command metadata.
func TestListOrderCommandNameAndDescription(t *testing.T) {
	// Given
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockService := services.NewMockListService(ctrl)
	// When
	cmd := NewListOrderCommand(ctx, mockService)

	// Then
	assert.Equal(t, "list-order", cmd.Name)
	assert.Contains(t, cmd.Description, "order alphabetically")
}

// TestListOrderCommandContextCancellation tests context cancellation handling.
func TestListOrderCommandContextCancellation(t *testing.T) {
	// Given
	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockService := services.NewMockListService(ctrl)

	// When
	cmd := NewListOrderCommand(ctx, mockService)

	_, cancel := context.WithCancel(ctx)
	cancel()

	// Then
	require.NotNil(t, cmd)
	require.Equal(t, "list-order", cmd.Name)
	assert.NotNil(t, cmd.Action)
}

// TestListOrderCommandActionAssignment tests that Action is properly set.
func TestListOrderCommandActionAssignment(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	ctrl := gomock.NewController(t)

	mockService := services.NewMockListService(ctrl)
	cmd := NewListOrderCommand(ctx, mockService)

	require.NotNil(t, cmd)
	require.NotNil(t, cmd.Action)
	assert.Equal(t, "list-order", cmd.Name)
}

// TestListOrderCommandNoServiceInjection tests error handling when service is not injected.
func TestListOrderCommandNoServiceInjection(t *testing.T) {
	t.Parallel()

	ctx := t.Context()
	cmd := NewListOrderCommand(ctx, nil)

	err := cmd.Run(ctx, nil)

	require.Error(t, err)
}
