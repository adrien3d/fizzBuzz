package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/adrien3d/fizzbuzz/fizzbuzz"
	"github.com/gin-gonic/gin"
)

// RequestStats stores the frequency of requests
type RequestStats struct {
	mu      sync.Mutex
	counter map[string]int
}

var stats = RequestStats{counter: make(map[string]int)}

func main() {
	r := gin.Default()
	r.GET("/fizzbuzz", fizzBuzzHandler)
	r.GET("/stats", statsHandler)

	_ = r.Run(":8080") // Start server on port 8080
}

func fizzBuzzHandler(c *gin.Context) {
	int1, err1 := strconv.Atoi(c.Query("int1"))
	int2, err2 := strconv.Atoi(c.Query("int2"))
	limit, err3 := strconv.Atoi(c.Query("limit"))
	str1 := c.Query("str1")
	str2 := c.Query("str2")

	if err1 != nil || err2 != nil || err3 != nil || int1 <= 0 || int2 <= 0 || limit <= 0 || str1 == "" || str2 == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters"})
		return
	}

	response := fizzbuzz.GenerateFizzBuzz(int1, int2, limit, str1, str2)

	// Track request
	stats.mu.Lock()
	key := fmt.Sprintf("%d-%d-%d-%s-%s", int1, int2, limit, str1, str2)
	stats.counter[key]++
	stats.mu.Unlock()

	c.JSON(http.StatusOK, response)
}

func statsHandler(c *gin.Context) {
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
