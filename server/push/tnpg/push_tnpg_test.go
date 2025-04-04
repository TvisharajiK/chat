package tnpg

import (
	"encoding/json"
	"testing"

	"net/http"
	"net/http/httptest"

	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinode/chat/server/push"
	"github.com/tinode/chat/server/store/types"
)

// Test generated using Keploy
func TestHandlerInit_ValidConfig_123(t *testing.T) {
	// Arrange
	validConfig := configType{
		Enabled:         true,
		OrgID:           "testOrg",
		AuthToken:       "testToken",
		DebugPushGWHost: "",
	}
	jsonConfig, err := json.Marshal(validConfig)
	require.NoError(t, err)

	// Act
	success, err := handler.Init(jsonConfig)

	// Assert
	require.NoError(t, err)
	assert.True(t, success)
	assert.Equal(t, "https://pushgw.tinode.co/pushv1/testorg", handler.pushUrl)
	assert.Equal(t, "https://pushgw.tinode.co/sub/testorg", handler.subUrl)
	assert.NotNil(t, handler.input)
	assert.NotNil(t, handler.channel)
	assert.NotNil(t, handler.stop)
}

// Test generated using Keploy
func TestHandlerInit_InvalidConfig_MissingOrgID_456(t *testing.T) {
	// Arrange
	invalidConfig := configType{
		Enabled:         true,
		OrgID:           "",
		AuthToken:       "testToken",
		DebugPushGWHost: "",
	}
	jsonConfig, err := json.Marshal(invalidConfig)
	require.NoError(t, err)

	// Act
	success, err := handler.Init(jsonConfig)

	// Assert
	require.Error(t, err)
	assert.False(t, success)
	assert.Equal(t, "organization name is missing", err.Error())
}

// Test generated using Keploy
func TestHandlerIsReady_Initialized_321(t *testing.T) {
	// Arrange
	handler.input = make(chan *push.Receipt)

	// Act
	ready := handler.IsReady()

	// Assert
	assert.True(t, ready)
}

// Test generated using Keploy
func TestHandlerStop_StopsCorrectly_654(t *testing.T) {
	// Arrange
	handler.stop = make(chan bool, 1)

	// Act
	handler.Stop()

	// Assert
	select {
	case stopped := <-handler.stop:
		assert.True(t, stopped)
	default:
		t.Fatal("Handler did not stop correctly")
	}
}

// Test generated using Keploy
func TestPostMessage_EncodeError_147(t *testing.T) {
	// Arrange
	// Channels cannot be JSON encoded, this will cause json.NewEncoder().Encode() to fail
	body := make(chan int)
	config := &configType{AuthToken: "testToken"}

	// Act
	resp, err := postMessage("http://dummy.url", body, config)

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "json: unsupported type: chan int")
	assert.Nil(t, resp)
}

// Test generated using Keploy
func TestPostMessage_NonGzipResponse_321(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// No Content-Encoding header
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&batchResponse{
			SuccessCount: 1,
			FailureCount: 0,
		})
	}))
	defer server.Close()
	originalClient := http.DefaultClient
	http.DefaultClient = server.Client()                      // Use test client
	t.Cleanup(func() { http.DefaultClient = originalClient }) // Restore default client

	body := map[string]string{"key": "value"}
	config := &configType{AuthToken: "testToken"}

	// Act
	resp, err := postMessage(server.URL, body, config)

	// Assert
	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, http.StatusOK, resp.httpCode)
	assert.Equal(t, 1, resp.SuccessCount)
	assert.Equal(t, 0, resp.FailureCount)
}

// Test generated using Keploy
func TestHandleSubResponse_NoFailures_010(t *testing.T) {
	// Arrange
	batch := &batchResponse{SuccessCount: 1, FailureCount: 0} // No failures
	req := &push.ChannelReq{Uid: types.ParseUid("usrSubOk")}
	devices := []string{"dSubOk"}
	channels := []string{}

	// Act
	handleSubResponse(batch, req, devices, channels) // Should return immediately

	// Assert
	// No external calls or panics expected. Log verification could ensure no warnings were logged.
}

// Test generated using Keploy
func TestHandlerInit_WithDebugHost_666(t *testing.T) {
	// Arrange
	debugHost := "http://localhost:18080" // Use a different port/path
	config := configType{
		Enabled:         true,
		OrgID:           "debugOrg",
		AuthToken:       "debugToken",
		DebugPushGWHost: debugHost,
	}
	jsonConfig, err := json.Marshal(config)
	require.NoError(t, err)
	originalHandler := handler

	// Act
	success, errInit := handler.Init(jsonConfig)

	// Assert
	require.NoError(t, errInit)
	assert.True(t, success)
	// OrgID is lowercased
	assert.Equal(t, "http://localhost:18080/pushv1/debugorg", handler.pushUrl)
	assert.Equal(t, "http://localhost:18080/sub/debugorg", handler.subUrl)
	assert.NotNil(t, handler.input)
	assert.NotNil(t, handler.channel)
	assert.NotNil(t, handler.stop)

	// Cleanup: Stop the worker goroutine if it was started
	if success && handler.stop != nil {
		handler.Stop()
		// Allow time for goroutine to exit - adjust as needed
		time.Sleep(50 * time.Millisecond)
	}
	handler = originalHandler // Restore if necessary
}
