package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestHub_topicDel_DeleteTopic_003(t *testing.T) {
	hub := &Hub{
		topics: &sync.Map{},
	}

	// Mock topic
	topicName := "testTopic"
	mockTopic := &Topic{name: topicName}
	hub.topics.Store(topicName, mockTopic)
	hub.numTopics = 1

	// Delete topic
	hub.topicDel(topicName)

	// Verify topic absence
	_, ok := hub.topics.Load(topicName)
	assert.False(t, ok)
	assert.Equal(t, 0, hub.numTopics)
}

// Test generated using Keploy
func TestHub_topicGet_ExistingAndNonExistingTopic_623(t *testing.T) {
	hub := &Hub{
		topics: &sync.Map{},
	}

	// Mock topic
	topicName := "testTopic"
	mockTopic := &Topic{name: topicName} // Use concrete Topic for simplicity in this test
	hub.topics.Store(topicName, mockTopic)

	// Test existing topic
	result := hub.topicGet(topicName)
	require.NotNil(t, result)
	assert.Equal(t, topicName, result.name) // Assuming Topic has a 'name' field

	// Test non-existing topic
	result = hub.topicGet("nonExistingTopic")
	assert.Nil(t, result)
}

// Test generated using Keploy
func TestHub_topicPut_AddTopic_511(t *testing.T) {
	hub := &Hub{
		topics: &sync.Map{},
	}

	// Mock topic
	topicName := "testTopic"
	mockTopic := &Topic{name: topicName} // Use concrete Topic

	// Add topic
	hub.topicPut(topicName, mockTopic)

	// Verify topic presence
	storedTopic, ok := hub.topics.Load(topicName)
	require.True(t, ok)
	assert.Equal(t, mockTopic, storedTopic)
	assert.Equal(t, 1, hub.numTopics)
}
