package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestGenDb_NoUsers_123(t *testing.T) {
	// Arrange
	data := &Data{
		Users: []User{},
	}

	// Act
	genDb(data)

	// Assert
	// Since no users are provided, the function should log "No data provided, stopping"
	// and return without performing any operations. No assertions are needed here
	// as the function does not return any value.
}

// Test generated using Keploy
func TestGetCreatedTime_ValidAndEmpty_A0A(t *testing.T) {
	now := time.Now().UTC().Round(time.Millisecond)

	// Valid duration
	delta := "-2h30m"
	expectedTime := now.Add(-2*time.Hour - 30*time.Minute)
	actualTime := getCreatedTime(delta)
	// Allow for slight difference due to execution time between now captures
	assert.WithinDuration(t, expectedTime, actualTime, 2*time.Millisecond)

	// Empty duration
	deltaEmpty := ""
	expectedTimeEmpty := now
	actualTimeEmpty := getCreatedTime(deltaEmpty)
	assert.WithinDuration(t, expectedTimeEmpty, actualTimeEmpty, 2*time.Millisecond)
}
