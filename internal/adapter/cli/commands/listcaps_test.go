package commands

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestListCapsCommandInitialization tests the initialization and basic behavior of ListCapsCommand
func TestListCapsCommandInitialization(t *testing.T) {
	// Given: A ListCapsCommand is initialized
	ctx := t.Context()
	cmd := NewListCapsCommand(ctx)

	// When: Run method is called with valid context
	err := cmd.Run(ctx, &cmd.Command)

	// Then: should return no error
	assert.NoError(t, err)
}

// TestListCapsCommandNameAndDescription tests the command metadata
func TestListCapsCommandNameAndDescription(t *testing.T) {
	// Given: A ListCapsCommand is created
	ctx := t.Context()
	cmd := NewListCapsCommand(ctx)

	// When: the command name and description are set correctly
	// Then: the Name should be "list-caps" and Description should contain "uppercase"
	assert.Equal(t, "list-caps", cmd.Name)
	assert.Contains(t, cmd.Description, "uppercase")
}

// TestListCapsCommandActionAssignment tests the Action field assignment
func TestListCapsCommandActionAssignment(t *testing.T) {
	// Given: A ListCapsCommand instance
	ctx := t.Context()
	cmd := NewListCapsCommand(ctx)

	// When: executing the command
	// Then: the Action field should be properly assigned
	require.NotNil(t, cmd)
	assert.NotNil(t, cmd.Action)
	require.Equal(t, "list-caps", cmd.Name)
}

// TestListOrderCommandExecution tests the command execution
func TestListCapsCommandExecution(t *testing.T) {
	// Given: A NewListCapsCommand instance with test data
	ctx := t.Context()
	cmd := NewListCapsCommand(ctx)

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

	assert.Contains(t, bufStr, "list-caps")
}
