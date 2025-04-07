package main

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestLpWrite_Success_104(t *testing.T) {
	wrt := httptest.NewRecorder()
	msg := []byte("hello world")

	err := lpWrite(wrt, msg)

	require.NoError(t, err)
	assert.Equal(t, "hello world", wrt.Body.String())
}
