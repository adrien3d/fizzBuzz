package main

import (
	"testing"
	"reflect"
)

func TestGenerateFizzBuzz(t *testing.T) {
	tests := []struct {
		int1, int2, limit int
		str1, str2       string
		expected         []string
	}{
		{3, 5, 15, "fizz", "buzz", []string{"1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz"}},
	}

	for _, tt := range tests {
		result := GenerateFizzBuzz(tt.int1, tt.int2, tt.limit, tt.str1, tt.str2)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("Expected %v, got %v", tt.expected, result)
		}
	}
}