package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tinode/chat/server/push"
	"github.com/tinode/chat/server/store/types"
)

// Test generated using Keploy
func TestPushForChanDelete_Basic_234(t *testing.T) {
	topicName := "grpToDelete234"
	now := types.TimeNow()
	expectedChannelName := types.GrpToChn(topicName)

	receipt := pushForChanDelete(topicName, now)
	require.NotNil(t, receipt)

	// Check Channel field
	assert.Equal(t, expectedChannelName, receipt.Channel)
	assert.Empty(t, receipt.To) // No individual recipients for channel delete

	// Check Payload
	assert.Equal(t, push.ActSub, receipt.Payload.What)
	assert.True(t, receipt.Payload.Silent)
	assert.Equal(t, expectedChannelName, receipt.Payload.Topic) // Topic in payload is channel name
	assert.Equal(t, now, receipt.Payload.Timestamp)
	assert.Equal(t, types.ModeNone, receipt.Payload.ModeWant)
	assert.Equal(t, types.ModeNone, receipt.Payload.ModeGiven)
	assert.Zero(t, receipt.Payload.SeqId)
	assert.Empty(t, receipt.Payload.From)
}
