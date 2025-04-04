package common

import (
	"testing"

	"fmt"
	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/tinode/chat/server/push"
	"google.golang.org/api/googleapi"
)

// Test generated using Keploy
func TestPayload_GetStringAttr_ValidField_001(t *testing.T) {
	payload := Payload{
		Title: "Test Title",
	}
	result := payload.getStringAttr("Title")
	assert.Equal(t, "Test Title", result)
}

// Test generated using Keploy
func TestPayload_GetStringAttr_InvalidField_002(t *testing.T) {
	payload := Payload{}
	result := payload.getStringAttr("NonExistentField")
	assert.Equal(t, "", result)
}

// Test generated using Keploy
func TestPayload_GetIntAttr_ValidField_003(t *testing.T) {
	payload := Payload{
		SummaryArgCount: 5,
	}
	result := payload.getIntAttr("SummaryArgCount")
	assert.Equal(t, 5, result)
}

// Test generated using Keploy
func TestPayload_GetIntAttr_InvalidField_004(t *testing.T) {
	payload := Payload{}
	result := payload.getIntAttr("NonExistentField")
	assert.Equal(t, 0, result)
}

// Test generated using Keploy
func TestConfig_GetStringField_MsgField_005(t *testing.T) {
	config := Config{
		Msg: Payload{
			Title: "Message Title",
		},
	}
	result := config.GetStringField(push.ActMsg, "Title")
	assert.Equal(t, "Message Title", result)
}

// Test generated using Keploy
func TestConfig_GetStringField_SubField_006(t *testing.T) {
	config := Config{
		Sub: Payload{
			Title: "Subscription Title",
		},
	}
	result := config.GetStringField(push.ActSub, "Title")
	assert.Equal(t, "Subscription Title", result)
}

// Test generated using Keploy
func TestConfig_GetStringField_CommonField_007(t *testing.T) {
	config := Config{
		Payload: Payload{
			Title: "Common Title",
		},
	}
	result := config.GetStringField(push.ActMsg, "Title")
	assert.Equal(t, "Common Title", result)
}

// Test generated using Keploy
func TestConfig_GetIntField_MsgField_008(t *testing.T) {
	config := Config{
		Msg: Payload{
			SummaryArgCount: 10,
		},
	}
	result := config.GetIntField(push.ActMsg, "SummaryArgCount")
	assert.Equal(t, 10, result)
}

// Test generated using Keploy
func TestConfig_GetIntField_SubField_009(t *testing.T) {
	config := Config{
		Sub: Payload{
			SummaryArgCount: 15,
		},
	}
	result := config.GetIntField(push.ActSub, "SummaryArgCount")
	assert.Equal(t, 15, result)
}

// Test generated using Keploy
func TestConfig_GetIntField_CommonField_010(t *testing.T) {
	config := Config{
		Payload: Payload{
			SummaryArgCount: 20,
		},
	}
	result := config.GetIntField(push.ActMsg, "SummaryArgCount")
	assert.Equal(t, 20, result)
}

// Test generated using Keploy
func TestDecodeGoogleApiError_GoogleApiError_011(t *testing.T) {
	googleErr := &googleapi.Error{
		Code:    400,
		Message: "Bad Request",
		Errors: []googleapi.ErrorItem{
			{Reason: "invalidArgument", Message: "Invalid argument provided"},
		},
		Details: []interface{}{
			map[string]interface{}{
				"@type":     "type.googleapis.com/google.firebase.fcm.v1.FcmError",
				"errorCode": "INVALID_ARGUMENT",
			},
		},
	}
	decoded, errs := DecodeGoogleApiError(googleErr)
	assert.Equal(t, 400, decoded.HttpCode)
	assert.Contains(t, decoded.ErrMessage, "Bad Request")
	assert.Contains(t, decoded.ErrMessage, "invalidArgument/Invalid argument provided")
	assert.Equal(t, "INVALID_ARGUMENT", decoded.FcmErrCode)
	assert.Empty(t, errs)
}

// Test generated using Keploy
func TestDecodeGoogleApiError_NonGoogleApiError_012(t *testing.T) {
	genericErr := fmt.Errorf("generic error")
	decoded, errs := DecodeGoogleApiError(genericErr)
	assert.Equal(t, http.StatusBadRequest, decoded.HttpCode)
	assert.Equal(t, "generic error", decoded.ErrMessage)
	assert.Equal(t, string(ErrorUnspecified), decoded.FcmErrCode)
	assert.Len(t, errs, 1)
	assert.Contains(t, errs[0].Error(), "not googleapi.Error")
}

// Test generated using Keploy
func TestPayload_GetStringAttr_NonStringField_518(t *testing.T) {
	// Payload where SummaryArgCount is an int, not a string.
	payload := Payload{
		SummaryArgCount: 5,
	}
	// Attempt to get SummaryArgCount as a string.
	result := payload.getStringAttr("SummaryArgCount")
	// Expect an empty string because the field is not of type string.
	assert.Equal(t, "", result)
}

// Test generated using Keploy
func TestPayload_GetIntAttr_NonIntField_391(t *testing.T) {
	// Payload where Title is a string, not an int.
	payload := Payload{
		Title: "Not an Integer",
	}
	// Attempt to get Title as an integer.
	result := payload.getIntAttr("Title")
	// Expect 0 because the field is not of type int.
	assert.Equal(t, 0, result)
}

// Test generated using Keploy
func TestDecodeGoogleApiError_GoogleApiError_NoDetails_654(t *testing.T) {
	googleErr := &googleapi.Error{
		Code:    500,
		Message: "Internal Server Error",
		Errors: []googleapi.ErrorItem{
			{Reason: "backendError", Message: "Something went wrong"},
		},
		Details: []interface{}{}, // Empty Details slice
	}
	decoded, errs := DecodeGoogleApiError(googleErr)
	assert.Equal(t, 500, decoded.HttpCode)
	assert.Contains(t, decoded.ErrMessage, "Internal Server Error")
	assert.Contains(t, decoded.ErrMessage, "backendError/Something went wrong")
	assert.Equal(t, ErrorUnspecified, decoded.FcmErrCode) // Default FcmErrCode
	assert.Empty(t, errs)
}

// Test generated using Keploy
func TestDecodeGoogleApiError_UnrecognizedDetailsFormat_812(t *testing.T) {
	googleErr := &googleapi.Error{
		Code:    400,
		Message: "Bad Request",
		Errors: []googleapi.ErrorItem{
			{Reason: "invalidArgument", Message: "Invalid argument provided"},
		},
		Details: []interface{}{
			"this is not a map", // Invalid detail format
			map[string]interface{}{ // A valid one might follow or not
				"@type":     "type.googleapis.com/google.firebase.fcm.v1.FcmError",
				"errorCode": "INVALID_ARGUMENT",
			},
		},
	}
	decoded, errs := DecodeGoogleApiError(googleErr)
	assert.Equal(t, 400, decoded.HttpCode)
	assert.Contains(t, decoded.ErrMessage, "Bad Request")
	assert.Contains(t, decoded.ErrMessage, "invalidArgument/Invalid argument provided")
	// FcmErrCode should be set from the valid detail item
	assert.Equal(t, ErrorInvalidArgument, decoded.FcmErrCode)
	// Check that an error was recorded for the invalid detail format
	assert.Len(t, errs, 1)
	assert.Contains(t, errs[0].Error(), "error.Details unrecognized format string")

	// Case 2: Only invalid details
	googleErrOnlyInvalid := &googleapi.Error{
		Code:    400,
		Message: "Bad Request",
		Details: []interface{}{"invalid detail"},
	}
	decoded2, errs2 := DecodeGoogleApiError(googleErrOnlyInvalid)
	assert.Equal(t, 400, decoded2.HttpCode)
	assert.Equal(t, ErrorUnspecified, decoded2.FcmErrCode) // Falls back to unspecified
	assert.Len(t, errs2, 1)
	assert.Contains(t, errs2[0].Error(), "error.Details unrecognized format string")
}
