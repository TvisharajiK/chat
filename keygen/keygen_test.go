package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestGenerate_WithHMACSalt_123(t *testing.T) {
	hmacSalt := "dGVzdGhhbWNzYWx0dGVzdGhhbWNzYWx0dGVzdGhhbWNzYWx0dGVzdGhhbWNzYWx0"
	sequence := 1
	isRoot := 0

	result := generate(sequence, isRoot, hmacSalt)

	assert.Equal(t, 0, result, "Expected generate to return 0")
}

// Test generated using Keploy
func TestGenerate_EmptyHMACSalt_456(t *testing.T) {
	hmacSalt := ""
	sequence := 1
	isRoot := 0

	result := generate(sequence, isRoot, hmacSalt)

	assert.Equal(t, 0, result, "Expected generate to return 0")
}
