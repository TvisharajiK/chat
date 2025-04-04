package push

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestRegister_NilHandler_002(t *testing.T) {
	// Arrange
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when registering nil handler")
		}
	}()

	// Act
	Register("nilHandler", nil)
}

// Test generated using Keploy
func TestInit_InvalidConfig_005(t *testing.T) {
	// Arrange
	invalidConfig := `invalid-json`

	// Act
	enabled, err := Init([]byte(invalidConfig))

	// Assert
	require.Error(t, err)
	assert.Nil(t, enabled)
	assert.Contains(t, err.Error(), "failed to parse config")
}

// Test generated using Keploy
func TestPush_NoHandlersNil_111(t *testing.T) {
	originalHandlers := handlers                   // Backup original
	handlers = nil                                 // Explicitly set to nil
	defer func() { handlers = originalHandlers }() // Restore

	// Arrange
	receipt := &Receipt{}

	// Act
	// Call Push and expect no panic or error
	assert.NotPanics(t, func() { Push(receipt) }, "Push should not panic with nil handlers")
}

// Test generated using Keploy
func TestChannelSub_NoHandlersNil_555(t *testing.T) {
	originalHandlers := handlers                   // Backup original
	handlers = nil                                 // Explicitly set to nil
	defer func() { handlers = originalHandlers }() // Restore

	// Arrange
	req := &ChannelReq{}

	// Act & Assert
	assert.NotPanics(t, func() { ChannelSub(req) }, "ChannelSub should not panic with nil handlers")
}

// Test generated using Keploy
func TestStop_NoHandlersNil_999(t *testing.T) {
	originalHandlers := handlers                   // Backup original
	handlers = nil                                 // Explicitly set to nil
	defer func() { handlers = originalHandlers }() // Restore

	// Act & Assert
	assert.NotPanics(t, func() { Stop() }, "Stop should not panic with nil handlers")
}
