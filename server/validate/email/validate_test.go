package email

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestValidatorInit_InvalidJSON_456(t *testing.T) {
	// Arrange
	invalidConfig := `{
        "host_url": "https://example.com",
        "languages": ["en", "es",
        "validation_templ": "{{.subject}}",
        "reset_secret_templ": "{{.body_plain}}",
        "sender": "sender@example.com",
        "login": "smtp-login",
        "sender_password": "smtp-password",
        "auth_mechanism": "plain",
        "max_retries": 5,
        "smtp_server": "smtp.example.com",
        "smtp_port": "587",
        "smtp_helo_host": "example.com",
        "insecure_skip_verify": false,
        "domains": ["example.com"],
        "code_length": 6
    }`

	validator := &validator{}

	// Act
	err := validator.Init(invalidConfig)

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "invalid character")
}

// Test generated using Keploy
func TestValidatorInit_UnknownAuthMechanism_234(t *testing.T) {
	// Arrange
	invalidAuthConfig := `{
        "host_url": "https://example.com",
        "languages": ["en"],
        "validation_templ": "{{.subject}}",
        "reset_secret_templ": "{{.body_plain}}",
        "sender": "sender@example.com",
        "login": "smtp-login",
        "sender_password": "smtp-password",
        "auth_mechanism": "unknown",
        "smtp_server": "smtp.example.com",
        "smtp_port": "587",
        "smtp_helo_host": "example.com"
    }`

	validator := &validator{}

	// Act
	err := validator.Init(invalidAuthConfig)

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unknown auth_mechanism")
}

// Test generated using Keploy
func TestValidatorInit_InvalidSenderFormat_112(t *testing.T) {
	// Arrange
	invalidSenderConfig := `{
        "sender": "invalid-email-format",
        "smtp_server": "smtp.example.com"
    }` // Minimal config to reach the sender parsing

	validator := &validator{}

	// Act
	err := validator.Init(invalidSenderConfig)

	// Assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "mail: missing '@' or angle-addr") // Specific error from mail.ParseAddress
}
