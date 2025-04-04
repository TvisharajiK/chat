package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinode/chat/server/store/types"
)

// Test generated using Keploy
func TestInitVideoCalls_ValidConfig_789(t *testing.T) {
	// Arrange
	jsconfig := json.RawMessage(`{
        "enabled": true,
        "call_establishment_timeout": 30,
        "ice_servers": [
            {"username": "user1", "credential": "pass1", "urls": ["stun:stun1.example.com"]}
        ]
    }`)

	// Act
	err := initVideoCalls(jsconfig)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, 30, globals.callEstablishmentTimeout)
	assert.Len(t, globals.iceServers, 1)
	assert.Equal(t, "user1", globals.iceServers[0].Username)
	assert.Equal(t, "pass1", globals.iceServers[0].Credential)
	assert.Equal(t, []string{"stun:stun1.example.com"}, globals.iceServers[0].Urls)
}

// Test generated using Keploy
func TestInitVideoCalls_InvalidConfig_321(t *testing.T) {
	// Arrange
	jsconfig := json.RawMessage(`{invalid-json}`)

	// Act
	err := initVideoCalls(jsconfig)

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to parse config")
}

// Test generated using Keploy
func TestVideoCallMessageHead_UpdateHeaders_654(t *testing.T) {
	// Arrange
	call := &videoCall{
		seq:         123,
		contentMime: "application/json",
	}
	head := map[string]any{
		"sender": "user123",
	}
	newState := "connected"
	duration := 60

	// Act
	updatedHead := call.messageHead(head, newState, duration)

	// Assert
	require.NotNil(t, updatedHead)
	assert.Equal(t, ":123", updatedHead["replace"])
	assert.Equal(t, "connected", updatedHead["webrtc"])
	assert.Equal(t, 60, updatedHead["webrtc-duration"])
	assert.Equal(t, "application/json", updatedHead["mime"])
	assert.Equal(t, "user123", updatedHead["sender"])
}

// Test generated using Keploy
func TestCallPartySession_ProxySession_987(t *testing.T) {
	// Arrange
	proxySession := &Session{
		proto:       PROXY,
		multi:       &Session{},
		sid:         "session-id",
		userAgent:   "test-agent",
		remoteAddr:  "127.0.0.1",
		lang:        "en",
		countryCode: "US",
		background:  true,
		uid:         types.ParseUid("12345"),
	}

	// Act
	result := callPartySession(proxySession)

	// Assert
	require.NotNil(t, result)
	assert.Equal(t, PROXY, result.proto)
	assert.Equal(t, proxySession.multi, result.multi)
	assert.Equal(t, "session-id", result.sid)
	assert.Equal(t, "test-agent", result.userAgent)
	assert.Equal(t, "127.0.0.1", result.remoteAddr)
	assert.Equal(t, "en", result.lang)
	assert.Equal(t, "US", result.countryCode)
	assert.True(t, result.background)
	assert.Equal(t, types.ParseUid("12345"), result.uid)
}

// Test generated using Keploy
func TestGetCallOriginator_NoCurrentCall_123(t *testing.T) {
	// Arrange
	topic := &Topic{
		currentCall: nil,
	}

	// Act
	uid, session := topic.getCallOriginator()

	// Assert
	assert.Equal(t, types.ZeroUid, uid)
	assert.Nil(t, session)
}

// Test generated using Keploy
func TestGetCallOriginator_WithOriginator_456(t *testing.T) {
	// Arrange
	originatorSession := &Session{}
	topic := &Topic{
		currentCall: &videoCall{
			parties: map[string]callPartyData{
				"originator-session": {
					uid:          types.ParseUid("12345"),
					isOriginator: true,
					sess:         originatorSession,
				},
			},
		},
	}

	// Act
	uid, session := topic.getCallOriginator()

	// Assert
	assert.Equal(t, types.ParseUid("12345"), uid)
	assert.Equal(t, originatorSession, session)
}

// Test generated using Keploy
func TestVideoCallInfoMessage_Basic_191(t *testing.T) {
	// Arrange
	call := &videoCall{
		seq: 999,
	}
	event := constCallEventOffer // Use a defined constant

	// Act
	msg := call.infoMessage(event)

	// Assert
	require.NotNil(t, msg)
	require.NotNil(t, msg.Info)
	assert.Equal(t, "call", msg.Info.What)
	assert.Equal(t, event, msg.Info.Event)
	assert.Equal(t, call.seq, msg.Info.SeqId)
	// Other fields should be zero/nil
	assert.Empty(t, msg.Info.Topic)
	assert.Empty(t, msg.Info.Src)
	assert.Empty(t, msg.Info.From)
	assert.Nil(t, msg.Info.Payload)
}

// Test generated using Keploy
func TestGetCallOriginator_NoOriginatorInParties_212(t *testing.T) {
	// Arrange
	nonOriginatorSession := &Session{}
	topic := &Topic{
		currentCall: &videoCall{
			parties: map[string]callPartyData{
				"non-originator-session": {
					uid:          types.ParseUid("usr54321"),
					isOriginator: false, // Explicitly false
					sess:         nonOriginatorSession,
				},
				"another-non-originator": {
					uid:          types.ParseUid("usr11223"),
					isOriginator: false,
					sess:         &Session{},
				},
			},
		},
	}

	// Act
	uid, session := topic.getCallOriginator()

	// Assert
	assert.Equal(t, types.ZeroUid, uid)
	assert.Nil(t, session)
}

// Test generated using Keploy
func TestVideoCallMessageHead_NilHead_172(t *testing.T) {
	// Arrange
	call := &videoCall{
		seq:         456,
		contentMime: "video/mp4",
	}
	var head map[string]any // nil map
	newState := "ringing"
	duration := 0 // Test zero duration as well

	// Act
	updatedHead := call.messageHead(head, newState, duration)

	// Assert
	require.NotNil(t, updatedHead, "Head should be initialized if nil")
	assert.Equal(t, ":456", updatedHead["replace"])
	assert.Equal(t, "ringing", updatedHead["webrtc"])
	assert.Equal(t, "video/mp4", updatedHead["mime"])
	_, durationExists := updatedHead["webrtc-duration"]
	assert.False(t, durationExists, "webrtc-duration should not exist for zero duration")
}
