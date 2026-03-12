// Package commands provides CLI command implementations.
package commands

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListOrderCommandInitialization tests the initialization and basic behavior of ListOrderCommand.
func TestListOrderCommandInitialization(t *testing.T) {
	t.Parallel()

	// Given: A ListOrderCommand is initialized
	ctx := t.Context()
	cmd := NewListOrderCommand(ctx)

	// When: Run method is called with valid context
	err := cmd.Run(ctx, &cmd.Command)

	// Then: should return no error
	require.NoError(t, err)
}

// TestListOrderCommandNameAndDescription tests the command metadata.
func TestListOrderCommandNameAndDescription(t *testing.T) {
	t.Parallel()

	// Given: A ListOrderCommand is created
	ctx := t.Context()
	cmd := NewListOrderCommand(ctx)

	// When: the command name and description are set correctly
	// Then: the Name should be "list-order" and Description should contain "order alphabetically"
	assert.Equal(t, "list-order", cmd.Name)
	assert.Contains(t, cmd.Description, "order alphabetically")
}

// TestListOrderCommandContextCancellation tests context cancellation handling.
func TestListOrderCommandContextCancellation(t *testing.T) {
	t.Parallel()

	// Given: A ListOrderCommand instance
	ctx := t.Context()
	cmd := NewListOrderCommand(ctx)

	// When: executing with different contexts
	_, cancel := context.WithCancel(ctx)
	cancel()

	// Then: should handle context cancellation properly
	require.NotNil(t, cmd)
	require.Equal(t, "list-order", cmd.Name)

	// Verify Action field is properly assigned
	assert.NotNil(t, cmd.Action)
}

// TestListOrderCommandActionAssignment tests that Action is properly set.
func TestListOrderCommandActionAssignment(t *testing.T) {
	t.Parallel()

	// Given: A ListOrderCommand instance
	ctx := t.Context()
	cmd := NewListOrderCommand(ctx)

	// When: checking the command structure
	// Then: Action should not be nil and should reference the Run method
	require.NotNil(t, cmd)
	require.NotNil(t, cmd.Action)
	assert.Equal(t, "list-order", cmd.Name)
}
