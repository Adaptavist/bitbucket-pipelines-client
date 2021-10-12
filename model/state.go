package model

import "strings"

// ResultError Container for error message
type ResultError struct {
	Message string `json:"message"`
}

// Result contains completion result of a pipeline
type Result struct {
	Name string `json:"name"`
	// Exists only when there is an error
	Error *ResultError `json:"error,omitempty"`
}

func (s Result) String() string {
	return s.Name
}

// HasError returns true is Error is not null
func (s Result) HasError() bool {
	return s.Error != nil
}

// OK if result name is successful.
func (s Result) OK() bool {
	return s.Name == "SUCCESSFUL"
}

// State of a pipeline or a step
type State struct {
	Name   string `json:"name"`
	Result Result `json:"result"`
}

func (s State) String() string {
	return strings.ToLower(s.Name)
}
