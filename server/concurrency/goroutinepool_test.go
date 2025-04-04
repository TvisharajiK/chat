package concurrency

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestNewGoRoutinePool_Creation_001(t *testing.T) {
	numWorkers := 5
	pool := NewGoRoutinePool(numWorkers)

	require.NotNil(t, pool)
	assert.NotNil(t, pool.work)
	assert.NotNil(t, pool.sem)
	assert.NotNil(t, pool.stop)
	assert.Equal(t, numWorkers, cap(pool.sem))
	assert.Equal(t, numWorkers, cap(pool.stop))
}

// Test generated using Keploy
func TestSchedule_TaskWithAvailableWorkers_002(t *testing.T) {
	numWorkers := 2
	pool := NewGoRoutinePool(numWorkers)

	taskExecuted := false
	task := func() {
		taskExecuted = true
	}

	pool.Schedule(task)

	// Allow some time for the task to execute
	time.Sleep(100 * time.Millisecond)

	assert.True(t, taskExecuted)
}

// Test generated using Keploy
func TestWorker_TaskProcessingAndStopSignal_005(t *testing.T) {
	numWorkers := 1
	pool := NewGoRoutinePool(numWorkers)

	taskExecuted := false
	task := func() {
		taskExecuted = true
	}

	go pool.worker(task)

	// Allow some time for the task to execute
	time.Sleep(100 * time.Millisecond)

	assert.True(t, taskExecuted)

	pool.Stop()

	// Allow some time for the stop signal to propagate
	time.Sleep(100 * time.Millisecond)

	// Ensure the worker has stopped
	assert.Equal(t, 0, len(pool.sem))
}
