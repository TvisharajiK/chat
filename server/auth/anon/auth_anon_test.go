package anon

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinode/chat/server/auth"
	"github.com/tinode/chat/server/store/types"
)

// Test generated using Keploy
func TestInit_BlankName_001(t *testing.T) {
	a := &authenticator{}
	err := a.Init(nil, "")
	require.Error(t, err)
	assert.Equal(t, "auth_anonymous: authenticator name cannot be blank", err.Error())
}

// Test generated using Keploy
func TestInit_AlreadyInitialized_002(t *testing.T) {
	a := &authenticator{name: "existing"}
	err := a.Init(nil, "newName")
	require.Error(t, err)
	assert.Equal(t, "auth_anonymous: already initialized as existing; newName", err.Error())
}

// Test generated using Keploy
func TestInit_Success_003(t *testing.T) {
	a := &authenticator{}
	err := a.Init(nil, "testName")
	require.NoError(t, err)
	assert.Equal(t, "testName", a.name)
}

// Test generated using Keploy
func TestIsInitialized_Initialized_004(t *testing.T) {
	a := &authenticator{name: "testName"}
	assert.True(t, a.IsInitialized())
}

// Test generated using Keploy
func TestAddRecord_LevelNone_006(t *testing.T) {
	rec := &auth.Rec{AuthLevel: auth.LevelNone}
	a := authenticator{}
	updatedRec, err := a.AddRecord(rec, nil, "127.0.0.1")
	require.NoError(t, err)
	assert.Equal(t, auth.LevelAnon, updatedRec.AuthLevel)
	assert.Equal(t, types.StateOK, updatedRec.State)
}

// Test generated using Keploy
func TestUpdateRecord_Success_008(t *testing.T) {
	rec := &auth.Rec{AuthLevel: auth.LevelAnon}
	a := authenticator{}
	updatedRec, err := a.UpdateRecord(rec, nil, "127.0.0.1")
	require.NoError(t, err)
	assert.Equal(t, rec, updatedRec)
}

// Test generated using Keploy
func TestAuthenticate_Unsupported_009(t *testing.T) {
	a := authenticator{}
	rec, secret, err := a.Authenticate(nil, "127.0.0.1")
	require.Error(t, err)
	assert.Nil(t, rec)
	assert.Nil(t, secret)
	assert.Equal(t, types.ErrUnsupported, err)
}

// Test generated using Keploy
func TestAsTag_EmptyString_010(t *testing.T) {
	a := authenticator{}
	tag := a.AsTag("someToken")
	assert.Equal(t, "", tag)
}

// Test generated using Keploy
func TestIsUnique_AlwaysTrue_011(t *testing.T) {
	a := authenticator{}
	unique, err := a.IsUnique(nil, "127.0.0.1")
	require.NoError(t, err)
	assert.True(t, unique)
}

// Test generated using Keploy
func TestGenSecret_Unsupported_012(t *testing.T) {
	a := authenticator{}
	secret, expiry, err := a.GenSecret(&auth.Rec{})
	require.Error(t, err)
	assert.Nil(t, secret)
	assert.Equal(t, time.Time{}, expiry)
	assert.Equal(t, types.ErrUnsupported, err)
}

// Test generated using Keploy
func TestDelRecords_Success_013(t *testing.T) {
	a := authenticator{}
	err := a.DelRecords(types.ZeroUid)
	require.NoError(t, err)
}

// Test generated using Keploy
func TestRestrictedTags_Nil_014(t *testing.T) {
	a := authenticator{}
	tags, err := a.RestrictedTags()
	require.NoError(t, err)
	assert.Nil(t, tags)
}

// Test generated using Keploy
func TestGetResetParams_Nil_015(t *testing.T) {
	a := authenticator{}
	params, err := a.GetResetParams(types.ZeroUid)
	require.NoError(t, err)
	assert.Nil(t, params)
}

// Test generated using Keploy
func TestGetRealName_HardcodedName_016(t *testing.T) {
	a := authenticator{}
	name := a.GetRealName()
	assert.Equal(t, "anonymous", name)
}
