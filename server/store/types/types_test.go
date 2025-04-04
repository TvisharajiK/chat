package types

import (
	"testing"

	"encoding/base64"
	"encoding/binary"

	"encoding/base32"
	"strings"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestStoreError_Error_123(t *testing.T) {
	// Arrange
	err := StoreError("test error")

	// Act
	result := err.Error()

	// Assert
	assert.Equal(t, "test error", result)
}

// Test generated using Keploy
func TestUid_IsZero_456(t *testing.T) {
	// Arrange
	var uid Uid = ZeroUid

	// Act
	result := uid.IsZero()

	// Assert
	assert.True(t, result)
}

// Test generated using Keploy
func TestUid_Compare_789(t *testing.T) {
	// Arrange
	uid1 := Uid(10)
	uid2 := Uid(20)
	uid3 := Uid(10)

	// Act & Assert
	assert.Equal(t, -1, uid1.Compare(uid2))
	assert.Equal(t, 1, uid2.Compare(uid1))
	assert.Equal(t, 0, uid1.Compare(uid3))
}

// Test generated using Keploy
func TestUid_MarshalBinary_321(t *testing.T) {
	// Arrange
	uid := Uid(12345)

	// Act
	result, err := uid.MarshalBinary()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, []byte{57, 48, 0, 0, 0, 0, 0, 0}, result)
}

// Test generated using Keploy
func TestUid_UnmarshalBinary_654(t *testing.T) {
	// Arrange
	var uid Uid
	data := []byte{57, 48, 0, 0, 0, 0, 0, 0}

	// Act
	err := uid.UnmarshalBinary(data)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, Uid(12345), uid)
}

// Test generated using Keploy
func TestUid_MarshalText_654(t *testing.T) {
	// Arrange
	zeroUid := Uid(0)
	nonZeroUid := Uid(12345)

	// Act & Assert for zero Uid
	result, err := zeroUid.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, []byte{}, result)

	// Act & Assert for non-zero Uid
	result, err = nonZeroUid.MarshalText()
	assert.NoError(t, err)
	expected := make([]byte, base64.URLEncoding.WithPadding(base64.NoPadding).EncodedLen(8))
	src := make([]byte, 8)
	binary.LittleEndian.PutUint64(src, uint64(nonZeroUid))
	base64.URLEncoding.WithPadding(base64.NoPadding).Encode(expected, src)
	assert.Equal(t, expected, result)
}

// Test generated using Keploy
func TestUid_String_321(t *testing.T) {
	// Arrange
	uid := Uid(12345)

	// Act
	result := uid.String()

	// Assert
	expected := make([]byte, base64.URLEncoding.WithPadding(base64.NoPadding).EncodedLen(8))
	src := make([]byte, 8)
	binary.LittleEndian.PutUint64(src, uint64(uid))
	base64.URLEncoding.WithPadding(base64.NoPadding).Encode(expected, src)
	assert.Equal(t, string(expected), result)
}

// Test generated using Keploy
func TestUid_String32_654(t *testing.T) {
	// Arrange
	uid := Uid(12345)

	// Act
	result := uid.String32()

	// Assert
	data := make([]byte, 8)
	binary.LittleEndian.PutUint64(data, uint64(uid))
	expected := strings.ToLower(base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(data))
	assert.Equal(t, expected, result)
}

// Test generated using Keploy
func TestParseUid_987(t *testing.T) {
	// Arrange
	validData := "AAAAAAAAAAA"
	invalidData := "invalid_base64"

	// Act & Assert for valid data
	result := ParseUid(validData)
	assert.Equal(t, Uid(0), result)

	// Act & Assert for invalid data
	result = ParseUid(invalidData)
	assert.Equal(t, Uid(0), result) // Default value since UnmarshalText fails silently
}

// Test generated using Keploy
func TestParseUid32_654(t *testing.T) {
	// Arrange
	validData := "aaaaaaaaaaaaa"
	invalidData := "invalid_base32"

	// Act & Assert for valid data
	result := ParseUid32(validData)
	assert.Equal(t, Uid(0), result)

	// Act & Assert for invalid data
	result = ParseUid32(invalidData)
	assert.Equal(t, Uid(0), result) // Default value since UnmarshalBinary fails silently
}

// Test generated using Keploy
func TestUid_UnmarshalBinary_InvalidLength_102(t *testing.T) {
	// Arrange
	var uid Uid
	data := []byte{1, 2, 3, 4} // Less than 8 bytes

	// Act
	err := uid.UnmarshalBinary(data)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "Uid.UnmarshalBinary: invalid length", err.Error())
	assert.Equal(t, Uid(0), uid) // Should remain unchanged (zero)
}

// Test generated using Keploy
func TestUid_UnmarshalText_DecodeError_104(t *testing.T) {
	// Arrange
	var uid Uid
	// Correct length (11) but invalid base64 characters
	invalidBase64Data := []byte("!!!!!!!!!!!")

	// Act
	err := uid.UnmarshalText(invalidBase64Data)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Uid.UnmarshalText: failed to decode")
	// Check if it contains the specific base64 error message
	assert.Contains(t, err.Error(), "illegal base64 data")
	assert.Equal(t, Uid(0), uid) // Should remain unchanged (zero)
}

// Test generated using Keploy
func TestUid_MarshalJSON_Zero_106(t *testing.T) {
	// Arrange
	uid := Uid(0)

	// Act
	jsonData, err := uid.MarshalJSON()

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, `""`, string(jsonData))
}

// Test generated using Keploy
func TestUid_UnmarshalJSON_ValidZero_108(t *testing.T) {
	// Arrange
	var uid Uid
	jsonData := []byte(`"AAAAAAAAAAA"`) // Represents Uid(0)

	// Act
	err := uid.UnmarshalJSON(jsonData)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, Uid(0), uid)
}

// Test generated using Keploy
func TestUid_UnmarshalJSON_InvalidLength_110(t *testing.T) {
	// Arrange
	var uid Uid
	invalidLengthJson := []byte(`"short"`) // Length != uidBase64Unpadded + 2

	// Act
	err := uid.UnmarshalJSON(invalidLengthJson)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Uid.UnmarshalJSON: invalid length")
	assert.Equal(t, Uid(0), uid)
}

// Test generated using Keploy
func TestUid_UnmarshalJSON_MissingQuotes_505(t *testing.T) {
	// Arrange
	var uid Uid
	// Use a valid base64 text of correct length, but without quotes
	baseUid := Uid(123)
	text, _ := baseUid.MarshalText() // Creates 11 byte valid base64 string
	missingQuotesJson := text

	// Act
	err := uid.UnmarshalJSON(missingQuotesJson)

	// Assert
	// The length check fails first because the length is 11, not 13
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Uid.UnmarshalJSON: invalid length")
	assert.Equal(t, Uid(0), uid)

	// Test specifically for the quote check (line 139)
	missingStartQuote := append([]byte{'a'}, text...)  // length 12
	missingStartQuote = append(missingStartQuote, '"') // length 13
	err = uid.UnmarshalJSON(missingStartQuote)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Uid.UnmarshalJSON: unrecognized")

	missingEndQuote := append([]byte{'"'}, text...) // length 12
	missingEndQuote = append(missingEndQuote, 'a')  // length 13
	err = uid.UnmarshalJSON(missingEndQuote)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Uid.UnmarshalJSON: unrecognized")
}

// Test generated using Keploy
func TestUid_UserId_509(t *testing.T) {
	// Arrange: Zero Uid
	zeroUid := ZeroUid

	// Act: Zero Uid
	resultZero := zeroUid.UserId()

	// Assert: Zero Uid
	assert.Equal(t, "", resultZero)

	// Arrange: Non-zero Uid
	nonZeroUid := Uid(7777777777777777777)
	expectedString := "usr" + nonZeroUid.String()

	// Act: Non-zero Uid
	resultNonZero := nonZeroUid.UserId()

	// Assert: Non-zero Uid
	assert.Equal(t, expectedString, resultNonZero)
}

// Test generated using Keploy
func TestUid_FndName_510(t *testing.T) {
	// Arrange: Zero Uid
	zeroUid := ZeroUid

	// Act: Zero Uid
	resultZero := zeroUid.FndName()

	// Assert: Zero Uid
	assert.Equal(t, "", resultZero)

	// Arrange: Non-zero Uid
	nonZeroUid := Uid(8888888888888888888)
	expectedString := "fnd" + nonZeroUid.String()

	// Act: Non-zero Uid
	resultNonZero := nonZeroUid.FndName()

	// Assert: Non-zero Uid
	assert.Equal(t, expectedString, resultNonZero)
}

// Test generated using Keploy
func TestParseUserId_Various_512(t *testing.T) {
	// Arrange: Valid case
	expectedUid := Uid(1010101010101010101)
	validIdString := expectedUid.UserId() // "usr" + base64 string

	// Act: Valid case
	resultValid := ParseUserId(validIdString)

	// Assert: Valid case
	assert.Equal(t, expectedUid, resultValid)

	// Arrange: Invalid prefix
	invalidPrefixString := "xxx" + expectedUid.String()

	// Act: Invalid prefix
	resultInvalidPrefix := ParseUserId(invalidPrefixString)

	// Assert: Invalid prefix
	assert.Equal(t, ZeroUid, resultInvalidPrefix)

	// Arrange: Invalid base64 content after prefix
	invalidBase64String := "usr!!!!!!!!!!!"

	// Act: Invalid base64 content
	resultInvalidBase64 := ParseUserId(invalidBase64String)

	// Assert: Invalid base64 content (UnmarshalText fails silently in ParseUserId)
	assert.Equal(t, ZeroUid, resultInvalidBase64)

	// Arrange: Empty string
	emptyString := ""

	// Act: Empty string
	resultEmpty := ParseUserId(emptyString)

	// Assert: Empty string
	assert.Equal(t, ZeroUid, resultEmpty)

	// Arrange: Prefix only
	prefixOnlyString := "usr"

	// Act: Prefix only
	resultPrefixOnly := ParseUserId(prefixOnlyString)

	// Assert: Prefix only (UnmarshalText fails)
	assert.Equal(t, ZeroUid, resultPrefixOnly)
}

// Test generated using Keploy
func TestGrpToChn_Various_513(t *testing.T) {
	assert.Equal(t, "chnABCDEFG", GrpToChn("grpABCDEFG"))
	assert.Equal(t, "chnABCDEFG", GrpToChn("chnABCDEFG")) // Already channel
	assert.Equal(t, "", GrpToChn("usrABCDEFG"))           // Different prefix
	assert.Equal(t, "", GrpToChn(""))                     // Empty string
	assert.Equal(t, "chn", GrpToChn("grp"))               // Prefix only
}

// Test generated using Keploy
func TestIsChannel_Various_514(t *testing.T) {
	assert.True(t, IsChannel("chnABCDEFG"))
	assert.True(t, IsChannel("chn")) // Prefix only
	assert.False(t, IsChannel("grpABCDEFG"))
	assert.False(t, IsChannel("usrABCDEFG"))
	assert.False(t, IsChannel(""))
	assert.False(t, IsChannel("ch")) // Incomplete prefix
}

// Test generated using Keploy
func TestChnToGrp_Various_515(t *testing.T) {
	assert.Equal(t, "grpABCDEFG", ChnToGrp("chnABCDEFG"))
	assert.Equal(t, "grpABCDEFG", ChnToGrp("grpABCDEFG")) // Already group
	assert.Equal(t, "", ChnToGrp("usrABCDEFG"))           // Different prefix
	assert.Equal(t, "", ChnToGrp(""))                     // Empty string
	assert.Equal(t, "grp", ChnToGrp("chn"))               // Prefix only
}
