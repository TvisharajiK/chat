package main

import (
	"net/http"
	"net/url"
	"testing"

	"io"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestServePprof_ValidPath_001(t *testing.T) {
	// Arrange
	mux := http.NewServeMux()
	serveAt := "debug/pprof"

	// Act
	servePprof(mux, serveAt)

	// Assert
	handler, pattern := mux.Handler(&http.Request{URL: &url.URL{Path: "/debug/pprof/"}})
	require.NotNil(t, handler)
	assert.Equal(t, "/debug/pprof/", pattern)
}

// Test generated using Keploy
func TestProfileHandler_UnknownProfile_003(t *testing.T) {
	// Arrange
	req := &http.Request{URL: &url.URL{Path: "/debug/pprof/unknown"}}
	wrt := httptest.NewRecorder()
	pprofHttpRoot = "/debug/pprof/"

	// Act
	profileHandler(wrt, req)

	// Assert
	result := wrt.Result()
	assert.Equal(t, http.StatusNotFound, result.StatusCode)
	body, _ := io.ReadAll(result.Body)
	assert.Contains(t, string(body), "Unknown profile 'unknown'")
}
