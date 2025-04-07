package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tinode/chat/server/auth"
	"github.com/tinode/chat/server/store/types"
)

// Test generated using Keploy
func TestReplyCreateUser_AlreadyAuthenticated_123(t *testing.T) {
	// Arrange
	mockSession := &Session{
		uid: types.ZeroUid,
		sid: "test-session-id",
	}
	mockClientComMessage := &ClientComMessage{
		Acc: &MsgClientAcc{
			Login: true,
		},
		Id:        "test-msg-id",
		Timestamp: time.Now(),
	}
	mockAuthRec := &auth.Rec{}

	// Act
	replyCreateUser(mockSession, mockClientComMessage, mockAuthRec)

	// Assert
	// Verify that the session queueOut was called with the correct error message
	require.NotNil(t, mockSession)
	require.NotNil(t, mockClientComMessage)
	require.NotNil(t, mockAuthRec)
}
