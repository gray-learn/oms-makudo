package main

import "context"

type store struct {
	// add here mongoDB
}

// constructor
func NewStore() *store { // return type is a pointer to store
	return &store{}
}

// Create inserts a new record into the store.
// It takes a context.Context parameter to allow for request-scoped values,
// deadlines, and cancellation signals.
// Returns an error if the operation fails.
func (s *store) Create(context.Context) error {
	return nil //
}
