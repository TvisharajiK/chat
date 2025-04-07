package main

import (
	"testing"
	"time"

	"errors"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestNewPromExporter_ValidInputs_123(t *testing.T) {
	// Arrange
	server := "http://localhost:8080"
	namespace := "tinode"
	timeout := 5 * time.Second
	scraper := &Scraper{}

	// Act
	exporter := NewPromExporter(server, namespace, timeout, scraper)

	// Assert
	require.NotNil(t, exporter)
	assert.Equal(t, server, exporter.address)
	assert.Equal(t, timeout, exporter.timeout)
	assert.Equal(t, namespace, exporter.namespace)
	assert.Equal(t, scraper, exporter.scraper)
	assert.NotNil(t, exporter.up)
	assert.NotNil(t, exporter.version)
	assert.NotNil(t, exporter.topicsLive)
	assert.NotNil(t, exporter.topicsTotal)
	assert.NotNil(t, exporter.sessionsLive)
	assert.NotNil(t, exporter.sessionsTotal)
	assert.NotNil(t, exporter.numGoroutines)
	assert.NotNil(t, exporter.incomingMessagesWebsockTotal)
	assert.NotNil(t, exporter.outgoingMessagesWebsockTotal)
	assert.NotNil(t, exporter.incomingMessagesLongpollTotal)
	assert.NotNil(t, exporter.outgoingMessagesLongpollTotal)
	assert.NotNil(t, exporter.incomingMessagesGrpcTotal)
	assert.NotNil(t, exporter.outgoingMessagesGrpcTotal)
	assert.NotNil(t, exporter.fileDownloadsTotal)
	assert.NotNil(t, exporter.fileUploadsTotal)
	assert.NotNil(t, exporter.ctrlCodesTotal2xx)
	assert.NotNil(t, exporter.ctrlCodesTotal3xx)
	assert.NotNil(t, exporter.ctrlCodesTotal4xx)
	assert.NotNil(t, exporter.ctrlCodesTotal5xx)
	assert.NotNil(t, exporter.clusterLeader)
	assert.NotNil(t, exporter.clusterSize)
	assert.NotNil(t, exporter.clusterNodesLive)
	assert.NotNil(t, exporter.malloced)
	assert.NotNil(t, exporter.requestLatencyMsCount)
	assert.NotNil(t, exporter.outgoingMessageBytesCount)
}

// Test generated using Keploy
func TestPromExporter_Describe_Valid_456(t *testing.T) {
	// Arrange
	server := "http://localhost:8080"
	namespace := "tinode"
	timeout := 5 * time.Second
	scraper := &Scraper{}
	exporter := NewPromExporter(server, namespace, timeout, scraper)
	ch := make(chan *prometheus.Desc, 100)

	// Act
	go func() {
		exporter.Describe(ch)
		close(ch)
	}()

	// Assert
	descs := []*prometheus.Desc{}
	for desc := range ch {
		descs = append(descs, desc)
	}
	require.NotEmpty(t, descs)
	assert.Contains(t, descs, exporter.up)
	assert.Contains(t, descs, exporter.version)
	assert.Contains(t, descs, exporter.topicsLive)
	assert.Contains(t, descs, exporter.topicsTotal)
	assert.Contains(t, descs, exporter.sessionsLive)
	assert.Contains(t, descs, exporter.sessionsTotal)
	assert.Contains(t, descs, exporter.numGoroutines)
	assert.Contains(t, descs, exporter.incomingMessagesWebsockTotal)
	assert.Contains(t, descs, exporter.outgoingMessagesWebsockTotal)
	assert.Contains(t, descs, exporter.incomingMessagesLongpollTotal)
	assert.Contains(t, descs, exporter.outgoingMessagesLongpollTotal)
	assert.Contains(t, descs, exporter.incomingMessagesGrpcTotal)
	assert.Contains(t, descs, exporter.outgoingMessagesGrpcTotal)
	assert.Contains(t, descs, exporter.fileDownloadsTotal)
	assert.Contains(t, descs, exporter.fileUploadsTotal)
	assert.Contains(t, descs, exporter.ctrlCodesTotal2xx)
	assert.Contains(t, descs, exporter.ctrlCodesTotal3xx)
	assert.Contains(t, descs, exporter.ctrlCodesTotal4xx)
	assert.Contains(t, descs, exporter.ctrlCodesTotal5xx)
	assert.Contains(t, descs, exporter.clusterLeader)
	assert.Contains(t, descs, exporter.clusterSize)
	assert.Contains(t, descs, exporter.clusterNodesLive)
	assert.Contains(t, descs, exporter.malloced)
	assert.Contains(t, descs, exporter.requestLatencyMsCount)
	assert.Contains(t, descs, exporter.outgoingMessageBytesCount)
}

// Test generated using Keploy
func TestFirstError_NoErrors_111(t *testing.T) {
	// Arrange
	errs := []error{nil, nil, nil}

	// Act
	result := firstError(errs...)

	// Assert
	assert.Nil(t, result)
}

// Test generated using Keploy
func TestFirstError_FirstError_222(t *testing.T) {
	// Arrange
	err1 := errors.New("first error")
	err2 := errors.New("second error")
	errs := []error{nil, err1, nil, err2}

	// Act
	result := firstError(errs...)

	// Assert
	require.Error(t, result)
	assert.Equal(t, err1, result)
}
