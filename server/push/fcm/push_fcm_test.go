package fcm

import (
	"testing"

	legacy "firebase.google.com/go/messaging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinode/chat/server/push"
	"github.com/tinode/chat/server/store/types"
)

// Test generated using Keploy
func TestHandler_Init_InvalidJSON_456(t *testing.T) {
	// Arrange
	invalidConfig := `{
        "enabled": true,
        "dry_run": false,
        "credentials": "{\"type\": \"service_account\", \"project_id\": \"test-project-id\"",
        "time_to_live": 3600
    ` // Missing closing braces

	// Act
	initialized, err := handler.Init([]byte(invalidConfig))

	// Assert
	require.Error(t, err)
	assert.False(t, initialized)
	assert.Contains(t, err.Error(), "failed to parse config")
}

// Test generated using Keploy
func TestHandler_IsReady_789(t *testing.T) {
	// Arrange
	handler.input = make(chan *push.Receipt, bufferSize)

	// Act
	ready := handler.IsReady()

	// Assert
	assert.True(t, ready)
}

// Test generated using Keploy
func TestHandler_Stop_012(t *testing.T) {
	// Arrange
	handler.stop = make(chan bool, 1)

	// Act
	handler.Stop()

	// Assert
	select {
	case <-handler.stop:
		assert.True(t, true)
	default:
		t.Fatal("Stop signal was not sent to the stop channel")
	}
}

// Test generated using Keploy
func TestHandler_Init_MissingCredentials_333(t *testing.T) {
	// Arrange
	// Reset handler state before test
	handler = Handler{}
	configJSON := `{"enabled": true}` // No credentials or credentials_file

	// Act
	initialized, err := handler.Init([]byte(configJSON))

	// Assert
	require.Error(t, err)
	assert.False(t, initialized)
	assert.EqualError(t, err, "missing credentials")
	assert.Nil(t, handler.input)
	assert.Nil(t, handler.channel)
	assert.Nil(t, handler.stop)
	assert.Nil(t, handler.client)
	assert.Nil(t, handler.v1)
	assert.Empty(t, handler.projectID)
}

// Test generated using Keploy
func TestHandler_Init_CredentialsFileError_555(t *testing.T) {
	// Arrange
	// Reset handler state before test
	handler = Handler{}
	// Use a non-existent file path
	configJSON := `{"enabled": true, "credentials_file": "/tmp/non-existent-fcm-creds-for-test.json"}`

	// Act
	initialized, err := handler.Init([]byte(configJSON))

	// Assert
	require.Error(t, err) // Expecting os.ReadFile error
	assert.False(t, initialized)
	// Error message depends on the OS, check for common pattern
	assert.Contains(t, err.Error(), "no such file or directory")
}

// Test generated using Keploy
func TestHandleSubErrors_NoFailures_111(t *testing.T) {
	// Arrange
	response := &legacy.TopicManagementResponse{
		FailureCount: 0, // No failures
		Errors:       nil,
	}
	uid := types.Uid(123)
	devices := []string{"device1"}

	// Act
	handleSubErrors(response, uid, devices)

	// Assert
	// No assertion needed, just ensure it doesn't panic or log errors incorrectly.
}
