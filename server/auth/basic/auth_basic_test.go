package basic

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinode/chat/server/auth"
	"github.com/tinode/chat/server/store/types"
)

// Test generated using Keploy
func TestInit_ValidConfig_123(t *testing.T) {
	// Arrange
	jsonConfig := json.RawMessage(`{
        "add_to_tags": true,
        "min_password_length": 5,
        "min_login_length": 3
    }`)
	authName := "test_authenticator"
	a := &authenticator{}

	// Act
	err := a.Init(jsonConfig, authName)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, authName, a.name)
	assert.Equal(t, true, a.addToTags)
	assert.Equal(t, 5, a.minPasswordLength)
	assert.Equal(t, 3, a.minLoginLength)
}

// Test generated using Keploy
func TestInit_InvalidName_456(t *testing.T) {
	// Arrange
	jsonConfig := json.RawMessage(`{
        "add_to_tags": true,
        "min_password_length": 5,
        "min_login_length": 3
    }`)
	authName := ""
	a := &authenticator{}

	// Act
	err := a.Init(jsonConfig, authName)

	// Assert
	require.Error(t, err)
	assert.Equal(t, "auth_basic: authenticator name cannot be blank", err.Error())
}

// Test generated using Keploy
func TestCheckLoginPolicy_ValidInvalidInputs_789(t *testing.T) {
	// Arrange
	a := &authenticator{
		minLoginLength: 3,
	}

	validLogin := "valid_login"
	invalidLogin := "x"

	// Act & Assert
	err := a.checkLoginPolicy(validLogin)
	require.NoError(t, err)

	err = a.checkLoginPolicy(invalidLogin)
	require.Error(t, err)
	assert.Equal(t, types.ErrPolicy, err)
}

// Test generated using Keploy
func TestCheckPasswordPolicy_ValidInvalidInputs_321(t *testing.T) {
	// Arrange
	a := &authenticator{
		minPasswordLength: 5,
	}

	validPassword := "valid_password"
	invalidPassword := "123"

	// Act & Assert
	err := a.checkPasswordPolicy(validPassword)
	require.NoError(t, err)

	err = a.checkPasswordPolicy(invalidPassword)
	require.Error(t, err)
	assert.Equal(t, types.ErrPolicy, err)
}

// Test generated using Keploy
func TestParseSecret_ValidInvalidInputs_654(t *testing.T) {
	// Arrange
	validSecret := []byte("username:password")
	invalidSecret := []byte("invalid_secret")

	// Act & Assert
	uname, password, err := parseSecret(validSecret)
	require.NoError(t, err)
	assert.Equal(t, "username", uname)
	assert.Equal(t, "password", password)

	uname, password, err = parseSecret(invalidSecret)
	require.Error(t, err)
	assert.Equal(t, types.ErrMalformed, err)
	assert.Empty(t, uname)
	assert.Empty(t, password)
}

// Test generated using Keploy
func TestIsInitialized_TrueFalse_987(t *testing.T) {
	// Arrange
	a := &authenticator{name: "test_authenticator"}
	b := &authenticator{}

	// Act & Assert
	assert.True(t, a.IsInitialized())
	assert.False(t, b.IsInitialized())
}

// Test generated using Keploy
func TestAddRecord_MalformedSecret_002(t *testing.T) {
	// Arrange
	rec := &auth.Rec{}
	secret := []byte("malformed_secret")
	remoteAddr := "127.0.0.1"

	a := &authenticator{
		name: "test_authenticator",
	}

	// Act
	result, err := a.AddRecord(rec, secret, remoteAddr)

	// Assert
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Equal(t, types.ErrMalformed, err)
}

// Test generated using Keploy
func TestAsTag_Variations_156(t *testing.T) {
	authName := "tag_test"
	aWithTags := &authenticator{
		name:              authName,
		addToTags:         true,
		minLoginLength:    3,
		minPasswordLength: 5,
	}
	aWithoutTags := &authenticator{
		name:              authName,
		addToTags:         false,
		minLoginLength:    3,
		minPasswordLength: 5,
	}

	// Test Tags Disabled
	t.Run("TagsDisabled", func(t *testing.T) {
		tag := aWithoutTags.AsTag("validlogin")
		assert.Empty(t, tag)
	})

	// Test Login Policy Fails
	t.Run("LoginPolicyFail", func(t *testing.T) {
		tag := aWithTags.AsTag("sh") // Too short
		assert.Empty(t, tag)
		tag = aWithTags.AsTag("invalid login") // Contains space
		assert.Empty(t, tag)
	})

	// Test Success
	t.Run("Success", func(t *testing.T) {
		token := "good.login_1"
		expectedTag := authName + ":" + token
		tag := aWithTags.AsTag(token)
		assert.Equal(t, expectedTag, tag)
	})
}

// Test generated using Keploy
func TestGenSecret_Unsupported_378(t *testing.T) {
	a := authenticator{}
	rec := &auth.Rec{Uid: types.ParseUid("usrGen123")}
	secret, expires, err := a.GenSecret(rec)
	assert.ErrorIs(t, err, types.ErrUnsupported)
	assert.Nil(t, secret)
	assert.True(t, expires.IsZero())
}

// Test generated using Keploy
func TestRestrictedTags_Variations_590(t *testing.T) {
	authName := "restrict_test"

	// Test Tags Enabled
	t.Run("TagsEnabled", func(t *testing.T) {
		a := &authenticator{name: authName, addToTags: true}
		prefixes, err := a.RestrictedTags()
		require.NoError(t, err)
		assert.Equal(t, []string{authName}, prefixes)
	})

	// Test Tags Disabled
	t.Run("TagsDisabled", func(t *testing.T) {
		a := &authenticator{name: authName, addToTags: false}
		prefixes, err := a.RestrictedTags()
		require.NoError(t, err)
		assert.Empty(t, prefixes)
	})
}

// Test generated using Keploy
func TestGetRealName_Constant_712(t *testing.T) {
	a := authenticator{}
	name := a.GetRealName()
	assert.Equal(t, realName, name)
	assert.Equal(t, "basic", name) // Verify the constant value
}
