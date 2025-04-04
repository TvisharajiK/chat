package mock_auth

import (
	"testing"

	"time"

	"encoding/json"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinode/chat/server/auth"
	"github.com/tinode/chat/server/store/types"
)

// Test generated using Keploy
func TestAddRecord_ValidInput_123(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		AddRecord(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&auth.Rec{Uid: types.Uid(12345)}, nil)

	rec := &auth.Rec{Uid: types.Uid(12345)}
	secret := []byte("secret")
	remoteAddr := "127.0.0.1"

	result, err := mockAuthHandler.AddRecord(rec, secret, remoteAddr)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, types.Uid(12345), result.Uid)
}

// Test generated using Keploy
func TestAsTag_ValidToken_456(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		AsTag("valid-token").
		Return("tag123")

	token := "valid-token"
	result := mockAuthHandler.AsTag(token)

	assert.Equal(t, "tag123", result)
}

// Test generated using Keploy
func TestAuthenticate_ValidInput_789(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		Authenticate(gomock.Any(), gomock.Any()).
		Return(&auth.Rec{Uid: types.Uid(12345)}, []byte("new-secret"), nil)

	secret := []byte("secret")
	remoteAddr := "127.0.0.1"

	rec, newSecret, err := mockAuthHandler.Authenticate(secret, remoteAddr)

	require.NoError(t, err)
	assert.NotNil(t, rec)
	assert.Equal(t, types.Uid(12345), rec.Uid)
	assert.Equal(t, []byte("new-secret"), newSecret)
}

// Test generated using Keploy
func TestDelRecords_ValidUid_321(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		DelRecords(types.Uid(12345)).
		Return(nil)

	uid := types.Uid(12345)
	err := mockAuthHandler.DelRecords(uid)

	require.NoError(t, err)
}

// Test generated using Keploy
func TestGenSecret_ValidRecord_654(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		GenSecret(gomock.Any()).
		Return([]byte("generated-secret"), time.Now().Add(time.Hour), nil)

	rec := &auth.Rec{Uid: types.Uid(12345)}

	secret, expiration, err := mockAuthHandler.GenSecret(rec)

	require.NoError(t, err)
	assert.NotNil(t, secret)
	assert.Equal(t, []byte("generated-secret"), secret)
	assert.True(t, expiration.After(time.Now()))
}

// Test generated using Keploy
func TestGetRealName_ValidResponse_111(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		GetRealName().
		Return("MockAuthHandler")

	result := mockAuthHandler.GetRealName()

	assert.Equal(t, "MockAuthHandler", result)
}

// Test generated using Keploy
func TestGetResetParams_ValidUid_222(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		GetResetParams(types.Uid(12345)).
		Return(map[string]interface{}{"param1": "value1"}, nil)

	uid := types.Uid(12345)
	result, err := mockAuthHandler.GetResetParams(uid)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "value1", result["param1"])
}

// Test generated using Keploy
func TestInit_ValidInput_333(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		Init(json.RawMessage(`{"key": "value"}`), "MockName").
		Return(nil)

	jsonConf := json.RawMessage(`{"key": "value"}`)
	name := "MockName"

	err := mockAuthHandler.Init(jsonConf, name)

	require.NoError(t, err)
}

// Test generated using Keploy
func TestIsInitialized_ValidResponse_444(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		IsInitialized().
		Return(true)

	result := mockAuthHandler.IsInitialized()

	assert.True(t, result)
}

// Test generated using Keploy
func TestIsUnique_ValidInput_555(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		IsUnique([]byte("unique-secret"), "127.0.0.1").
		Return(true, nil)

	secret := []byte("unique-secret")
	remoteAddr := "127.0.0.1"

	result, err := mockAuthHandler.IsUnique(secret, remoteAddr)

	require.NoError(t, err)
	assert.True(t, result)
}

// Test generated using Keploy
func TestRestrictedTags_ValidResponse_666(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		RestrictedTags().
		Return([]string{"tag1", "tag2"}, nil)

	result, err := mockAuthHandler.RestrictedTags()

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, []string{"tag1", "tag2"}, result)
}

// Test generated using Keploy
func TestUpdateRecord_ValidInput_777(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthHandler := NewMockAuthHandler(ctrl)
	mockAuthHandler.EXPECT().
		UpdateRecord(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(&auth.Rec{Uid: types.Uid(12345)}, nil)

	rec := &auth.Rec{Uid: types.Uid(12345)}
	secret := []byte("updated-secret")
	remoteAddr := "127.0.0.1"

	result, err := mockAuthHandler.UpdateRecord(rec, secret, remoteAddr)

	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, types.Uid(12345), result.Uid)
}
