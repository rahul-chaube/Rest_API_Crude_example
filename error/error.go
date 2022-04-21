package main

type StatusError struct {
	ErrorCode int `json:"error_code"`
	Err       error
}
