package models

import "errors"

// FizzBuzzRequest gathers the parameters of a fizzbuzz request
type FizzBuzzRequest struct {
	Int1  int    `form:"int1"`
	Int2  int    `form:"int2"`
	Limit int    `form:"limit"`
	Str1  string `form:"str1"`
	Str2  string `form:"str2"`
}

func (req FizzBuzzRequest) Validate() error {
	if req.Int1 <= 0 {
		return errors.New("int1 must be a strictly positive integer")
	}
	if req.Int2 <= 0 {
		return errors.New("int2 must be a strictly positive integer")
	} 
	if req.Limit <= 0 {
		return errors.New("limit must be a strictly positive integer")
	} 
	if req.Str1 == "" {
		return errors.New("str1 must be a not empty string")
	} 
	if req.Str2 == "" {
		return errors.New("str2 must be a not empty string")
	}
	return nil
}

// FizzBuzzResponse is the response for a fizzbuzz request
type FizzBuzzResponse struct {
	Result []string `json:"result"`
}