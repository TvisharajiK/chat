package main

import (
	"testing"

	"bytes"
	"log"
	"os"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestParseMetricList_ValidInput_123(t *testing.T) {
	input := "metric1, metric2 ,metric3"
	expected := []string{"metric1", "metric2", "metric3"}

	result := parseMetricList(input)

	assert.Equal(t, expected, result, "The parsed metric list should match the expected output")
}

// Test generated using Keploy
func TestPromHTTPLogger_Println_512(t *testing.T) {
	// Capture log output
	var buf bytes.Buffer
	log.SetOutput(&buf)
	originalFlags := log.Flags()
	log.SetFlags(0) // Disable timestamp etc.
	defer func() {
		log.SetOutput(os.Stderr)
		log.SetFlags(originalFlags)
	}()

	logger := promHTTPLogger{}
	logger.Println("hello", "world", 123)

	assert.Equal(t, "hello world 123\n", buf.String())
}
