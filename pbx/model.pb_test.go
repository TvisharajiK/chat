package pbx

import (

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Test generated using Keploy
func TestAuthLevelEnum_ValidValue_001(t *testing.T) {
	authLevel := AuthLevel_AUTH
	enumPtr := authLevel.Enum()
	require.NotNil(t, enumPtr)
	assert.Equal(t, AuthLevel_AUTH, *enumPtr)
}

// Test generated using Keploy
func TestAuthLevelString_ValidValue_002(t *testing.T) {
	authLevel := AuthLevel_ROOT
	str := authLevel.String()
	assert.Equal(t, "ROOT", str)
}

// Test generated using Keploy
func TestInfoNoteEnum_ValidValue_003(t *testing.T) {
	infoNote := InfoNote_RECV
	enumPtr := infoNote.Enum()
	require.NotNil(t, enumPtr)
	assert.Equal(t, InfoNote_RECV, *enumPtr)
}

// Test generated using Keploy
func TestInfoNoteString_ValidValue_004(t *testing.T) {
	infoNote := InfoNote_CALL
	str := infoNote.String()
	assert.Equal(t, "CALL", str)
}

// Test generated using Keploy
func TestCallEventEnum_ValidValue_005(t *testing.T) {
	callEvent := CallEvent_INVITE
	enumPtr := callEvent.Enum()
	require.NotNil(t, enumPtr)
	assert.Equal(t, CallEvent_INVITE, *enumPtr)
}

// Test generated using Keploy
func TestCallEventString_ValidValue_006(t *testing.T) {
	callEvent := CallEvent_RINGING
	str := callEvent.String()
	assert.Equal(t, "RINGING", str)
}

// Test generated using Keploy
func TestAuthLevelEnumDescriptor_ValidValue_007(t *testing.T) {
	descriptor, indexes := AuthLevel(0).EnumDescriptor()
	require.NotNil(t, descriptor)
	require.NotNil(t, indexes)
	assert.Equal(t, []int{0}, indexes)
}

// Test generated using Keploy
func TestInfoNoteEnumDescriptor_ValidValue_008(t *testing.T) {
	descriptor, indexes := InfoNote(0).EnumDescriptor()
	require.NotNil(t, descriptor)
	require.NotNil(t, indexes)
	assert.Equal(t, []int{1}, indexes)
}

// Test generated using Keploy
func TestCallEventEnumDescriptor_ValidValue_009(t *testing.T) {
	descriptor, indexes := CallEvent(0).EnumDescriptor()
	require.NotNil(t, descriptor)
	require.NotNil(t, indexes)
	assert.Equal(t, []int{2}, indexes)
}

// Test generated using Keploy
func TestAuthLevelNumber_ValidValue_010(t *testing.T) {
	authLevel := AuthLevel_AUTH
	number := authLevel.Number()
	assert.Equal(t, protoreflect.EnumNumber(20), number)
}

// Test generated using Keploy
func TestInfoNoteNumber_ValidValue_011(t *testing.T) {
	infoNote := InfoNote_RECV
	number := infoNote.Number()
	assert.Equal(t, protoreflect.EnumNumber(2), number)
}

// Test generated using Keploy
func TestCallEventNumber_ValidValue_012(t *testing.T) {
	callEvent := CallEvent_RINGING
	number := callEvent.Number()
	assert.Equal(t, protoreflect.EnumNumber(7), number)
}

// Test generated using Keploy
func TestAuthLevelType_456(t *testing.T) {
	// Arrange
	authLevel := AuthLevel_NONE

	// Act
	enumType := authLevel.Type()

	// Assert
	require.NotNil(t, enumType)
	assert.Equal(t, protoreflect.FullName("pbx.AuthLevel"), enumType.Descriptor().FullName())
}

// Test generated using Keploy
func TestInfoNoteType_101(t *testing.T) {
	// Arrange
	infoNote := InfoNote_READ

	// Act
	enumType := infoNote.Type()

	// Assert
	require.NotNil(t, enumType)
	assert.Equal(t, protoreflect.FullName("pbx.InfoNote"), enumType.Descriptor().FullName())
}

// Test generated using Keploy
func TestCallEventType_131(t *testing.T) {
	// Arrange
	callEvent := CallEvent_HANG_UP

	// Act
	enumType := callEvent.Type()

	// Assert
	require.NotNil(t, enumType)
	assert.Equal(t, protoreflect.FullName("pbx.CallEvent"), enumType.Descriptor().FullName())
}
