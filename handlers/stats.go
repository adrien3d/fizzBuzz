package handlers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// RequestStats stores the frequency of requests
type RequestStats struct {
	mu      sync.Mutex
	counter map[string]int
}

var stats = RequestStats{counter: make(map[string]int)}

func ResetStatsForTest() {
	stats.mu.Lock()
	defer stats.mu.Unlock()
	stats.counter = make(map[string]int)
}

func RecordTestStat(key string, count int) {
	stats.mu.Lock()
	defer stats.mu.Unlock()
	stats.counter[key] = count
}

func StatsHandler(c *gin.Context) {
	stats.mu.Lock()
	var maxKey string
	var maxCount int
	for k, v := range stats.counter {
		if v > maxCount {
			maxKey = k
			maxCount = v
		}
	}
	stats.mu.Unlock()

	if maxKey == "" {
		c.JSON(http.StatusOK, gin.H{"message": "No requests recorded yet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"most_frequent_request": maxKey, "hits": maxCount})
}