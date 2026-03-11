package util

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ToJSON_SimpleObject(t *testing.T) {
	// Given
	src := map[string]string{"key": "value"}

	// When
	result := ToJSON(src)

	// Then
	if result != `{"key":"value"}` {
		t.Errorf("Expected %v, got %v", `{"key":"value"}`, result)
	}
}

func Test_ToJSON_NestedObject(t *testing.T) {
	// Given
	src := map[string]interface{}{"outer": map[string]string{"inner": "data"}}

	// When
	result := ToJSON(src)

	// Then
	if result != `{"outer":{"inner":"data"}}` {
		t.Errorf("Expected %v, got %v", `{"outer":{"inner":"data"}}`, result)
	}
}

func Test_ToJSON_EmptyObject(t *testing.T) {
	// Given
	src := map[string]string{}

	// When
	result := ToJSON(src)

	// Then
	if result != `{}` {
		t.Errorf("Expected %v, got %v", `{}`, result)
	}
}

func Test_ToJSON_InvalidJSONString(t *testing.T) {
	// Given
	src := math.Inf(1)

	// When
	result := ToJSON(src)

	// Then
	assert.Empty(t, result)
}
