// Package commands provides CLI command implementations.
package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListCapsCommandInitialization tests the initialization and basic behavior of ListCapsCommand.
func TestListCapsCommandInitialization(t *testing.T) {
	t.Parallel()

	// Given: A ListCapsCommand is initialized
	ctx := t.Context()
	cmd := NewListCapsCommand(ctx)

	// When: Run method is called with valid context
	err := cmd.Run(ctx, &cmd.Command)

	// Then: should return no error
	require.NoError(t, err)
}

// TestListCapsCommandNameAndDescription tests the command metadata.
func TestListCapsCommandNameAndDescription(t *testing.T) {
	t.Parallel()

	// Given: A ListCapsCommand is created
	ctx := t.Context()
	cmd := NewListCapsCommand(ctx)

	// When: the command name and description are set correctly
	// Then: the Name should be "list-caps" and Description should contain "uppercase"
	assert.Equal(t, "list-caps", cmd.Name)
	assert.Contains(t, cmd.Description, "uppercase")
}

// TestListCapsCommandActionAssignment tests that Action is properly set.
func TestListCapsCommandActionAssignment(t *testing.T) {
	t.Parallel()

	// Given: A ListCapsCommand instance
	ctx := t.Context()
	cmd := NewListCapsCommand(ctx)

	// When: checking the command structure
	// Then: the Action field should be properly assigned and the name should be correct
	require.NotNil(t, cmd)
	assert.NotNil(t, cmd.Action)
	require.Equal(t, "list-caps", cmd.Name)
}

// TestListCapsCommandExecution tests the command execution.
func TestListCapsCommandExecution(t *testing.T) {
	t.Parallel()

	// Given: A ListCapsCommand instance
	ctx := t.Context()
	cmd := NewListCapsCommand(ctx)

	// When: the Action method is invoked
	err := cmd.Action(ctx, &cmd.Command)

	// Then: it should execute without error
	require.NoError(t, err)
	assert.Equal(t, "list-caps", cmd.Name)
}
