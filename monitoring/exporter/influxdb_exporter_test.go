package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestFormAuthorizationHeaderValue_Version2_0_112(t *testing.T) {
	token := "my-secret-token-v2"
	version := "2.0"

	headerValue := formAuthorizationHeaderValue(version, token)

	expected := "Token my-secret-token-v2"
	assert.Equal(t, expected, headerValue)
}

// Test generated using Keploy
func TestFormAuthorizationHeaderValue_Version1_7_131(t *testing.T) {
	token := "my-secret-token-v17"
	version := "1.7"

	headerValue := formAuthorizationHeaderValue(version, token)

	expected := "Bearer my-secret-token-v17"
	assert.Equal(t, expected, headerValue)
}
