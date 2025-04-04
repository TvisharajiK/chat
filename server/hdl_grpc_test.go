package main

import (
	"testing"

	"sync"

	"crypto/tls"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	pbxMocks "github.com/tinode/chat/mocks/pbx"
	"github.com/tinode/chat/pbx"
)

// Test generated using Keploy
func TestGrpcWrite_Success_004(t *testing.T) {
	// Arrange
	mockStream := new(pbxMocks.Node_MessageLoopServer)
	mockSession := &Session{
		grpcnode: mockStream,
	}
	testMessage := &pbx.ServerMsg{}

	mockStream.On("Send", testMessage).Return(nil)

	// Act
	err := grpcWrite(mockSession, testMessage)

	// Assert
	require.NoError(t, err)
	mockStream.AssertCalled(t, "Send", testMessage)
}

// Test generated using Keploy
func TestCloseGrpc_IsGRPC_519(t *testing.T) {
	// Arrange
	mockStream := new(pbxMocks.Node_MessageLoopServer)
	sess := &Session{
		proto:    GRPC,
		grpcnode: mockStream,
		lock:     sync.Mutex{},
	}

	// Act
	sess.closeGrpc()

	// Assert
	assert.Nil(t, sess.grpcnode, "grpcnode should be nil after closeGrpc for GRPC protocol")
}

// Test generated using Keploy
func TestGrpcWrite_NodeNil_794(t *testing.T) {
	// Arrange
	// No need for mockStream here as grpcnode is nil
	sess := &Session{
		grpcnode: nil, // Explicitly set to nil
	}
	testMessage := &pbx.ServerMsg{}

	// Act
	err := grpcWrite(sess, testMessage)

	// Assert
	require.NoError(t, err, "grpcWrite should return nil when grpcnode is nil")
	// No mock stream to assert calls on, the absence of panic/error is the assertion
}

// Test generated using Keploy
func TestServeGrpc_NoAddr_127(t *testing.T) {
	// Arrange
	addr := ""
	tlsConf := &tls.Config{}
	kaEnabled := true

	// Act
	server, err := serveGrpc(addr, kaEnabled, tlsConf)

	// Assert
	require.NoError(t, err, "serveGrpc should not return error for empty address")
	require.Nil(t, server, "serveGrpc should return nil server for empty address")
}
