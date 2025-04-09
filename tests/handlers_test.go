package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/adrien3d/fizzbuzz/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzHandler_ValidRequest(t *testing.T) {
	// Setup router and test context
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/fizzbuzz", handlers.FizzBuzzHandler)
	// Send a valid query
	req, _ := http.NewRequest("GET", "/fizzbuzz?int1=3&int2=5&limit=5&str1=Fizz&str2=Buzz", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var result []string
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, []string{"1", "2", "Fizz", "4", "Buzz"}, result)
}

func TestFizzBuzzHandler_InvalidQuery(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/fizzbuzz", handlers.FizzBuzzHandler)

	// Missing parameters (invalid query)
	req, _ := http.NewRequest("GET", "/fizzbuzz?int1=3", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	var body map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &body)
	assert.NoError(t, err)
	assert.Contains(t, body["error"], "Invalid")
}


// Mock or reset the stats before each test
func resetStats() {
	// Exported StatsHandler uses internal `stats` variable, so we must ensure it's reset for test isolation
	handlers.ResetStatsForTest()
}

func TestStatsHandler_NoRequests(t *testing.T) {
	resetStats()

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/stats", handlers.StatsHandler)

	req, _ := http.NewRequest("GET", "/stats", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "No requests recorded yet")
}

func TestStatsHandler_WithRequests(t *testing.T) {
	resetStats()

	// Simulate some stats
	handlers.RecordTestStat("3-5-100-Fizz-Buzz", 10)
	handlers.RecordTestStat("3-5-15-Fizz-Buzz", 5)

	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/stats", handlers.StatsHandler)

	req, _ := http.NewRequest("GET", "/stats", nil)
	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "3-5-100-Fizz-Buzz")
	assert.Contains(t, resp.Body.String(), `"hits":10`)
}