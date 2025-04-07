package fcm

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinode/chat/server/push"
	"github.com/tinode/chat/server/push/common"
)

// Test generated using Keploy
func TestPayloadToData_NilPayload_456(t *testing.T) {
	// Act
	data, err := payloadToData(nil)

	// Assert
	require.Error(t, err)
	assert.Nil(t, data)
	assert.EqualError(t, err, "empty push payload")
}

// Test generated using Keploy
func TestClonePayload_DeepCopy_789(t *testing.T) {
	// Arrange
	original := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	// Act
	cloned := clonePayload(original)

	// Assert
	require.NotNil(t, cloned)
	assert.Equal(t, original, cloned)

	// Modify the cloned map and ensure the original is unaffected
	cloned["key1"] = "modified_value"
	assert.NotEqual(t, original["key1"], cloned["key1"])
}

// Test generated using Keploy
func TestPayloadToData_UnknownPushType_456(t *testing.T) {
	// Arrange
	payload := &push.Payload{
		What: "unknown_type",
	}

	// Act
	data, err := payloadToData(payload)

	// Assert
	require.Error(t, err)
	assert.Nil(t, data)
	assert.EqualError(t, err, "unknown push type")
}

// Test generated using Keploy
func TestPayloadToData_ReadAction_106(t *testing.T) {
	// Arrange
	timestamp := time.Now()
	payload := &push.Payload{
		What:      push.ActRead,
		Topic:     "usr456",
		Timestamp: timestamp,
		From:      "sys", // Or the user sending the read notification
		SeqId:     15,
	}

	// Act
	data, err := payloadToData(payload)

	// Assert
	require.NoError(t, err)
	require.NotNil(t, data)
	assert.Equal(t, push.ActRead, data["what"])
	assert.Equal(t, "usr456", data["topic"])
	assert.Equal(t, timestamp.Format(time.RFC3339Nano), data["ts"])
	assert.Equal(t, "sys", data["xfrom"])
	assert.Equal(t, "15", data["seq"])
	assert.Equal(t, "true", data["silent"]) // Read notifications are always silent (line 84)
}

// Test generated using Keploy
func TestPrepareV1Notifications_NoRecipientsOrChannel_209(t *testing.T) {
	// Arrange
	rcpt := &push.Receipt{
		To:      nil, // No recipients
		Channel: "",  // No channel
		Payload: push.Payload{What: push.ActMsg, Topic: "grp123"},
	}
	config := &configType{}
	// No mocks needed as store won't be called (line 114 check fails early)

	// Act
	messages, uids := PrepareV1Notifications(rcpt, config)

	// Assert
	assert.Nil(t, messages) // Should return nil (line 133-135)
	assert.Nil(t, uids)
}

// Test generated using Keploy
func TestAndroidNotificationConfig_MsgWithConfig_501(t *testing.T) {
	// Arrange
	what := push.ActMsg
	topic := "grp123"
	data := map[string]string{"content": "message content"} // Data needed for $content
	config := &configType{
		TimeToLive: 7200,
		Android: &common.Config{
			Enabled: true, // Enable notification part
			Payload: common.Payload{
				// Define fields used in lines 286-299
				TitleLocKey: "title_key",
				Title:       "New Message Title",
				BodyLocKey:  "body_key",
				Body:        "$content", // Use content from data
				Icon:        "msg_icon",
				Color:       "#FF0000",
				ClickAction: "OPEN_CHAT",
			},
		},
	}

	// Act
	ac := androidNotificationConfig(what, topic, data, config)

	// Assert
	require.NotNil(t, ac)
	assert.Equal(t, string(common.AndroidPriorityHigh), ac.Priority)                                      // Line 262
	assert.Equal(t, "7200s", ac.Ttl)                                                                      // Line 245
	require.NotNil(t, ac.Notification)                                                                    // Line 286
	assert.Equal(t, topic, ac.Notification.Tag)                                                           // Line 289
	assert.Equal(t, string(common.AndroidNotificationPriorityHigh), ac.Notification.NotificationPriority) // Line 281
	assert.Equal(t, string(common.AndroidVisibilityPrivate), ac.Notification.Visibility)                  // Line 291
	assert.Equal(t, "title_key", ac.Notification.TitleLocKey)                                             // Line 292
	assert.Equal(t, "New Message Title", ac.Notification.Title)                                           // Line 293
	assert.Equal(t, "body_key", ac.Notification.BodyLocKey)                                               // Line 294
	assert.Equal(t, "message content", ac.Notification.Body)                                              // Check $content replacement (line 276-278, 295)
	assert.Equal(t, "msg_icon", ac.Notification.Icon)                                                     // Line 296
	assert.Equal(t, "#FF0000", ac.Notification.Color)                                                     // Line 297
	assert.Equal(t, "OPEN_CHAT", ac.Notification.ClickAction)                                             // Line 298
}

// Test generated using Keploy
func TestAndroidNotificationConfig_Read_502(t *testing.T) {
	// Arrange
	what := push.ActRead
	topic := "usr456"
	data := map[string]string{}            // Content doesn't matter for read
	config := &configType{TimeToLive: 100} // Custom TTL

	// Act
	ac := androidNotificationConfig(what, topic, data, config) // Enters line 248 branch

	// Assert
	require.NotNil(t, ac)
	assert.Equal(t, string(common.AndroidPriorityNormal), ac.Priority) // Line 250
	assert.Equal(t, "100s", ac.Ttl)                                    // Line 252 (uses config TTL)
	assert.Nil(t, ac.Notification)                                     // Line 251
}

// Test generated using Keploy
func TestAndroidNotificationConfig_WebRTC_503(t *testing.T) {
	// Arrange
	what := push.ActMsg
	topic := "grp789"
	data := map[string]string{"webrtc": "started", "content": "video call"} // Line 256 checks data["webrtc"]
	config := &configType{
		TimeToLive: 3600, // Default TTL, should be overridden
		Android: &common.Config{
			Enabled: true, // Enable notification part
			Payload: common.Payload{Title: "Incoming Call", Body: "$content"},
		},
	}

	// Act
	ac := androidNotificationConfig(what, topic, data, config)

	// Assert
	require.NotNil(t, ac)
	assert.Equal(t, string(common.AndroidPriorityHigh), ac.Priority) // Line 262
	assert.Equal(t, "0s", ac.Ttl)                                    // Line 258 sets TTL to 0s for video call
	require.NotNil(t, ac.Notification)                               // Line 286
	assert.Equal(t, topic, ac.Notification.Tag)                      // Line 289
	// Priority should be max for video calls (line 282-284)
	assert.Equal(t, string(common.AndroidNotificationPriorityMax), ac.Notification.NotificationPriority) // Line 290
	assert.Equal(t, "Incoming Call", ac.Notification.Title)                                              // Line 293
	assert.Equal(t, "video call", ac.Notification.Body)                                                  // Body should still be set (line 295)
}

// Test generated using Keploy
func TestAndroidNotificationConfig_Disabled_504(t *testing.T) {
	// Arrange
	what := push.ActMsg
	topic := "grp123"
	data := map[string]string{"content": "message content"}
	config := &configType{
		TimeToLive: 3600,
		Android:    &common.Config{Enabled: false}, // Android disabled (line 271 check)
	}

	// Act
	ac := androidNotificationConfig(what, topic, data, config)

	// Assert
	require.NotNil(t, ac)
	assert.Equal(t, string(common.AndroidPriorityHigh), ac.Priority) // Priority still High
	assert.Equal(t, "3600s", ac.Ttl)                                 // TTL still set
	assert.Nil(t, ac.Notification)                                   // Notification part should be nil (returns at line 272)
}

// Test generated using Keploy
func TestApnsNotificationConfig_MsgWithAlert_701(t *testing.T) {
	// Arrange
	what := push.ActMsg
	topic := "grptopic"
	data := map[string]string{"content": "apns message"} // For $content replacement
	unread := 3
	bundleId := "com.example.app"
	config := &configType{
		TimeToLive:   120, // Custom TTL
		ApnsBundleID: bundleId,
		Apns: &common.Config{
			Enabled: true, // Enable alert generation
			Payload: common.Payload{
				// Fields used in lines 350-361
				Action:          "ACT",
				ActionLocKey:    "ACT_KEY",
				Body:            "$content", // Use content from data
				LaunchImage:     "img.png",
				LocKey:          "MSG_KEY",
				Title:           "APNS Title",
				Subtitle:        "APNS Subtitle",
				TitleLocKey:     "TITLE_KEY",
				SummaryArg:      "sender_name",
				SummaryArgCount: 1,
			},
		},
	}

	// Act
	ac := apnsNotificationConfig(what, topic, data, unread, config) // Enters line 344 block

	// Assert
	require.NotNil(t, ac)
	require.NotNil(t, ac.Headers) // Line 368
	require.NotNil(t, ac.Payload) // Line 378

	// Check Headers (Lines 369-374)
	assert.Equal(t, bundleId, ac.Headers[common.HeaderApnsTopic])
	assert.Equal(t, topic, ac.Headers[common.HeaderApnsCollapseID])
	assert.Equal(t, "10", ac.Headers[common.HeaderApnsPriority])                             // Default high priority (line 316)
	assert.Equal(t, string(common.ApnsPushTypeAlert), ac.Headers[common.HeaderApnsPushType]) // Default alert type (line 315)
	expiresStr := ac.Headers[common.HeaderApnsExpiration]
	expiresUnix, err := strconv.ParseInt(expiresStr, 10, 64)
	require.NoError(t, err)
	// Check expiration uses config.TimeToLive (line 311-313)
	expectedExpires := time.Now().UTC().Add(120 * time.Second).Unix()
	assert.InDelta(t, expectedExpires, expiresUnix, 5, "Expiration time mismatch")

	// Check Payload (Lines 364-367 marshal map[string]interface{}{"aps": apsPayload})
	var payload map[string]interface{}
	err = json.Unmarshal(ac.Payload, &payload)
	require.NoError(t, err)
	aps, ok := payload["aps"].(map[string]interface{})
	require.True(t, ok)

	// Check APS fields (Lines 334-341)
	assert.Equal(t, float64(unread), aps["badge"])
	assert.Equal(t, float64(1), aps["content-available"])
	assert.Equal(t, float64(1), aps["mutable-content"])
	assert.Equal(t, string(common.InterruptionLevelTimeSensitive), aps["interruption-level"]) // Default for msg (line 317)
	assert.Equal(t, "default", aps["sound"])
	assert.Equal(t, topic, aps["thread-id"])

	// Check Alert fields (Lines 350-361)
	alert, ok := aps["alert"].(map[string]interface{})
	require.True(t, ok, "Alert should be present")
	assert.Equal(t, config.Apns.GetStringField(what, "Action"), alert["action"])
	assert.Equal(t, config.Apns.GetStringField(what, "ActionLocKey"), alert["action-loc-key"])
	assert.Equal(t, "apns message", alert["body"]) // Check $content replacement (lines 346-348, 353)
	assert.Equal(t, config.Apns.GetStringField(what, "LaunchImage"), alert["launch-image"])
	assert.Equal(t, config.Apns.GetStringField(what, "LocKey"), alert["loc-key"])
	assert.Equal(t, config.Apns.GetStringField(what, "Title"), alert["title"])
	assert.Equal(t, config.Apns.GetStringField(what, "Subtitle"), alert["subtitle"])
	assert.Equal(t, config.Apns.GetStringField(what, "TitleLocKey"), alert["title-loc-key"])
	assert.Equal(t, config.Apns.GetStringField(what, "SummaryArg"), alert["summary-arg"])
	assert.Equal(t, float64(config.Apns.GetIntField(what, "SummaryArgCount")), alert["summary-arg-count"])
}

// Test generated using Keploy
func TestApnsNotificationConfig_WebRTCStarted_703(t *testing.T) {
	// Arrange
	what := push.ActMsg
	topic := "p2pabc"
	data := map[string]string{"webrtc": "started"} // Call started (line 309, 318 check)
	unread := 1
	bundleId := "com.example.voip"
	config := &configType{
		// TimeToLive is ignored for VOIP, defaultTimeToLive used initially then overridden
		ApnsBundleID: bundleId,
		Apns:         &common.Config{Enabled: true}, // Enabled, but call shouldn't alert visually (line 305 check)
	}

	// Act
	ac := apnsNotificationConfig(what, topic, data, unread, config)

	// Assert
	require.NotNil(t, ac)
	require.NotNil(t, ac.Headers)
	require.NotNil(t, ac.Payload)

	// Check Headers
	// NOTE: Currently, code comments out the .voip suffix and voip push type (lines 321-326).
	// Test reflects the current *actual* behavior ('alert' type, normal bundleId).
	assert.Equal(t, bundleId, ac.Headers[common.HeaderApnsTopic]) // Uses normal bundleId (line 314, not 326)
	assert.Equal(t, topic, ac.Headers[common.HeaderApnsCollapseID])
	assert.Equal(t, "10", ac.Headers[common.HeaderApnsPriority])                             // High priority (line 316)
	assert.Equal(t, string(common.ApnsPushTypeAlert), ac.Headers[common.HeaderApnsPushType]) // Uses 'alert', not 'voip' (line 315, not 325)

	// Check VOIP expiration (short TTL, line 327)
	expiresStr := ac.Headers[common.HeaderApnsExpiration]
	expiresUnix, err := strconv.ParseInt(expiresStr, 10, 64)
	require.NoError(t, err)
	expectedExpires := time.Now().UTC().Add(voipTimeToLive * time.Second).Unix()
	assert.InDelta(t, expectedExpires, expiresUnix, 5)

	// Check Payload
	var payload map[string]interface{}
	err = json.Unmarshal(ac.Payload, &payload)
	require.NoError(t, err)
	aps, ok := payload["aps"].(map[string]interface{})
	require.True(t, ok)
	assert.Equal(t, float64(unread), aps["badge"])
	assert.Equal(t, float64(1), aps["content-available"])
	assert.Equal(t, string(common.InterruptionLevelCritical), aps["interruption-level"]) // Critical for calls (line 320)
	_, alertPresent := aps["alert"]
	assert.False(t, alertPresent, "Alert should not be present for WebRTC calls") // Checked by apnsShouldPresentAlert (line 344)
}
