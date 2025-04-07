package main

import (
	"testing"

	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test generated using Keploy
func TestParseMetric_Success_591(t *testing.T) {
	// Arrange
	stats := map[string]interface{}{
		"level1": map[string]interface{}{
			"level2": 123.45,
		},
	}
	path := "level1.level2"

	// Act
	value, err := parseMetric(stats, path)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, 123.45, value)
}

// Test generated using Keploy
func TestParseMetric_InvalidPath_IntermediateNotFound_333(t *testing.T) {
	// Arrange
	stats := map[string]interface{}{
		"level1": map[string]interface{}{
			"anotherKey": 123.45,
		},
	}
	path := "level1.level2" // level2 does not exist

	// Act
	value, err := parseMetric(stats, path)

	// Assert
	require.Error(t, err)
	assert.Equal(t, errKeyNotFound, err)
	assert.Equal(t, 0, value) // Default return on error
}

// Test generated using Keploy
func TestParseMetric_InvalidPath_NotAMap_472(t *testing.T) {
	// Arrange
	stats := map[string]interface{}{
		"level1": 123.45, // level1 is not a map
	}
	path := "level1.level2"

	// Act
	value, err := parseMetric(stats, path)

	// Assert
	require.Error(t, err)
	assert.Equal(t, errKeyNotFound, err)
	assert.Equal(t, 0, value) // Default return on error
}

// Test generated using Keploy
func TestParseNumeric_Success_118(t *testing.T) {
	// Arrange
	stats := map[string]interface{}{
		"metric": 123.45,
	}
	path := "metric"

	// Act
	value, err := parseNumeric(stats, path)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, 123.45, value)
}

// Test generated using Keploy
func TestParseNumeric_ParseMetricError_921(t *testing.T) {
	// Arrange
	stats := map[string]interface{}{}
	path := "nonexistent.metric"

	// Act
	value, err := parseNumeric(stats, path)

	// Assert
	require.Error(t, err)
	assert.Equal(t, errKeyNotFound, err)
	assert.Equal(t, 0.0, value)
}

// Test generated using Keploy
func TestParseNumeric_NotAFloat_882(t *testing.T) {
	// Arrange
	stats := map[string]interface{}{
		"metric": "not a float",
	}
	path := "metric"

	// Act
	value, err := parseNumeric(stats, path)

	// Assert
	require.Error(t, err)
	assert.Equal(t, errKeyNotFound, err) // Returns errKeyNotFound in this case
	assert.Equal(t, 0.0, value)
}

// Test generated using Keploy
func TestParseList_Success_401(t *testing.T) {
	// Arrange
	stats := map[string]interface{}{
		"list": []interface{}{1.1, 2.2, 3.3},
	}
	path := "list"
	expected := []float64{1.1, 2.2, 3.3}

	// Act
	value, err := parseList(stats, path)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, expected, value)
}

// Test generated using Keploy
func TestParseList_ParseMetricError_511(t *testing.T) {
	// Arrange
	stats := map[string]interface{}{}
	path := "nonexistent.list"

	// Act
	value, err := parseList(stats, path)

	// Assert
	require.Error(t, err)
	assert.Equal(t, errKeyNotFound, err)
	assert.Nil(t, value)
}

// Test generated using Keploy
func TestParseList_NotASlice_215(t *testing.T) {
	// Arrange
	stats := map[string]interface{}{
		"list": "not a slice",
	}
	path := "list"

	// Act
	value, err := parseList(stats, path)

	// Assert
	require.Error(t, err)
	assert.Equal(t, errMalformed, err)
	assert.Nil(t, value)
}

// Test generated using Keploy
func TestParseHisto_CountParseError_802(t *testing.T) {
	// Arrange
	stats := map[string]interface{}{
		"histo.sum":              100.0,
		"histo.count_per_bucket": []interface{}{1.0, 2.0, 7.0},
		"histo.bounds":           []interface{}{10.0, 20.0},
	} // Missing histo.count
	key := "histo"

	// Act
	h, err := parseHisto(stats, key)

	// Assert
	require.Error(t, err)
	assert.Equal(t, errKeyNotFound, err)
	assert.Nil(t, h)
}

// Test generated using Keploy
func TestParseStatsRaw_SimpleMetricError_838(t *testing.T) {
	// Arrange
	scraper := &Scraper{
		simpleMetrics: []string{"simple1", "missing"}, // "missing" key does not exist
	}
	stats := map[string]interface{}{
		"simple1": 1.0,
	}

	// Act
	metrics, err := scraper.parseStatsRaw(stats)

	// Assert
	require.Error(t, err)
	assert.Equal(t, errKeyNotFound, err)
	assert.Nil(t, metrics)
}

// Test generated using Keploy
func TestParseStatsRaw_HistogramMetricError_912(t *testing.T) {
	// Arrange
	scraper := &Scraper{
		simpleMetrics:    []string{"simple1"},
		histogramMetrics: []string{"histo1"}, // histo1.count is missing
	}
	stats := map[string]interface{}{
		"simple1": 1.0,
		// Missing histo1.count
		"histo1.sum":              50.0,
		"histo1.count_per_bucket": []interface{}{1.0, 4.0},
		"histo1.bounds":           []interface{}{10.0},
	}

	// Act
	metrics, err := scraper.parseStatsRaw(stats)

	// Assert
	require.Error(t, err)
	assert.Equal(t, errKeyNotFound, err)
	assert.Nil(t, metrics)
}

// Test generated using Keploy
func TestScrape_Success_145(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"metric": 123.45}`)
	}))
	defer server.Close()

	scraper := &Scraper{address: server.URL}

	// Act
	stats, err := scraper.Scrape()

	// Assert
	require.NoError(t, err)
	require.NotNil(t, stats)
	assert.Equal(t, 123.45, stats["metric"])
}

// Test generated using Keploy
func TestScrape_HTTPGetError_664(t *testing.T) {
	// Arrange
	// Intentionally use an invalid address to cause a connection error
	scraper := &Scraper{address: "http://invalid-address-that-does-not-exist:12345"}

	// Act
	stats, err := scraper.Scrape()

	// Assert
	require.Error(t, err) // Error should be connection refused or similar
	assert.Nil(t, stats)
}

// Test generated using Keploy
func TestCollectRaw_ScrapeError_693(t *testing.T) {
	// Arrange
	scraper := &Scraper{address: "http://invalid-address-that-does-not-exist:12345"}

	// Act
	metrics, err := scraper.CollectRaw()

	// Assert
	require.Error(t, err)
	// Check if the error is network-related (optional, depends on exact error wrapping)
	// assert.Contains(t, err.Error(), "connect: connection refused") // Example check
	assert.Nil(t, metrics)
	// Check that 'up' metric is not added on error
	_, exists := metrics["up"]
	assert.False(t, exists)
}

// Test generated using Keploy
func TestCollectRaw_ParseStatsError_448(t *testing.T) {
	// Arrange
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		// Missing "simple" metric required by the scraper config
		fmt.Fprintln(w, `{}`)
	}))
	defer server.Close()

	scraper := &Scraper{
		address:          server.URL,
		simpleMetrics:    []string{"simple"}, // Requires "simple" key
		histogramMetrics: []string{},
	}

	// Act
	metrics, err := scraper.CollectRaw()

	// Assert
	require.Error(t, err)
	assert.Equal(t, errKeyNotFound, err) // Error originates from parseNumeric -> parseMetric
	assert.Nil(t, metrics)
	// Check that 'up' metric is not added on error
	_, exists := metrics["up"]
	assert.False(t, exists)
}

// Test generated using Keploy
func TestParseStatsRaw_EmptyMetrics_789(t *testing.T) {
	// Arrange
	scraper := &Scraper{
		simpleMetrics:    []string{}, // Empty list
		histogramMetrics: []string{}, // Empty list
	}
	stats := map[string]interface{}{
		"some_other_data": 123.0, // Data exists but isn't requested
	}

	// Act
	metrics, err := scraper.parseStatsRaw(stats)

	// Assert
	require.NoError(t, err, "Parsing with empty metric lists should not error")
	require.NotNil(t, metrics, "Metrics map should be initialized")
	assert.Empty(t, metrics, "Metrics map should be empty as no metrics were requested")
}
