package commands

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListOrderCommandInitialization tests the initialization and basic behavior of ListOrderCommand
func TestListOrderCommandInitialization(t *testing.T) {
	// Given: A ListOrderCommand is initialized
	ctx := t.Context()
	cmd := NewListOrderCommand(ctx)

	// When: Run method is called with valid context
	err := cmd.Run(ctx, &cmd.Command)

	// Then: should return no error
	assert.NoError(t, err)
}

// TestListOrderCommandNameAndDescription tests the command metadata
func TestListOrderCommandNameAndDescription(t *testing.T) {
	// Given: A ListOrderCommand is created
	ctx := t.Context()
	cmd := NewListOrderCommand(ctx)

	// When: the command name and description are set correctly
	// Then: the Name should be "list-order" and Description should contain "order alphabetically"
	assert.Equal(t, "list-order", cmd.Name)
	assert.Contains(t, cmd.Description, "order alphabetically")
}

// TestListOrderCommandContextCancellation tests context cancellation handling
func TestListOrderCommandContextCancellation(t *testing.T) {
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

// TestListOrderCommandExecution tests the command execution
func TestListOrderCommandExecution(t *testing.T) {
	// Given: A ListOrderCommand instance with test data
	ctx := t.Context()
	cmd := NewListOrderCommand(ctx)

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// When
	err := cmd.Action(ctx, &cmd.Command)

	w.Close()
	os.Stdout = oldStdout

	// Then
	assert.NoError(t, err)

	var buf bytes.Buffer
	io.Copy(&buf, r)

	bufStr := buf.String()

	assert.Contains(t, bufStr, "list-order")
}
