package tests

import (
	"errors"
	"testing"

	"github.com/adrien3d/fizzbuzz/models"
	"github.com/stretchr/testify/assert"
)

func TestFizzBuzzRequest_Validate(t *testing.T) {
	tests := []struct {
		name    string
		req     models.FizzBuzzRequest
		wantErr error
	}{
		{
			name: "valid request",
			req:  models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 100, Str1: "Fizz", Str2: "Buzz"},
			wantErr: nil,
		},
		{
			name:    "int1 is zero",
			req:     models.FizzBuzzRequest{Int1: 0, Int2: 5, Limit: 100, Str1: "Fizz", Str2: "Buzz"},
			wantErr: errors.New("int1 must be a strictly positive integer"),
		},
		{
			name:    "int2 is negative",
			req:     models.FizzBuzzRequest{Int1: 3, Int2: -2, Limit: 100, Str1: "Fizz", Str2: "Buzz"},
			wantErr: errors.New("int2 must be a strictly positive integer"),
		},
		{
			name:    "limit is zero",
			req:     models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 0, Str1: "Fizz", Str2: "Buzz"},
			wantErr: errors.New("limit must be a strictly positive integer"),
		},
		{
			name:    "str1 is empty",
			req:     models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 100, Str1: "", Str2: "Buzz"},
			wantErr: errors.New("str1 must be a not empty string"),
		},
		{
			name:    "str2 is empty",
			req:     models.FizzBuzzRequest{Int1: 3, Int2: 5, Limit: 100, Str1: "Fizz", Str2: ""},
			wantErr: errors.New("str2 must be a not empty string"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.req.Validate()

			if tt.wantErr == nil {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tt.wantErr.Error())
			}
		})
	}
}