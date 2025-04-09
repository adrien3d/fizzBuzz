package handlers

import (
	"fmt"
	"net/http"

	"github.com/adrien3d/fizzbuzz/models"
	"github.com/adrien3d/fizzbuzz/services"
	"github.com/gin-gonic/gin"
)

func FizzBuzzHandler(c *gin.Context) {
	var req models.FizzBuzzRequest
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameters:" + err.Error()})
		return
	}

	response := services.GenerateFizzBuzz(req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)

	// Track request
	stats.mu.Lock()
	key := fmt.Sprintf("%d-%d-%d-%s-%s", req.Int1, req.Int2, req.Limit, req.Str1, req.Str2)
	stats.counter[key]++
	stats.mu.Unlock()

	c.JSON(http.StatusOK, response)
}