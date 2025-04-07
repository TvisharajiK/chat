package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	storeMocks "github.com/tinode/chat/mocks/server/store"
	"github.com/tinode/chat/server/store"
)

// Test generated using Keploy
func TestLargeFileRunGarbageCollection_PeriodBlockSize_004(t *testing.T) {
	// Arrange
	mockFilePersistence := new(storeMocks.FilePersistenceInterface)
	store.Files = mockFilePersistence

	period := time.Second * 10
	blockSize := 5
	stopChan := largeFileRunGarbageCollection(period, blockSize)

	mockFilePersistence.On("DeleteUnused", mock.Anything, blockSize).Return(nil)

	// Act
	time.Sleep(time.Second * 12) // Allow GC to run at least once
	stopChan <- true

	// Assert
	mockFilePersistence.AssertExpectations(t)
}
