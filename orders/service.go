package main

import "context"

type service struct {
	store OrderStore
}

func NewService(store OrderStore) *service { // return type is a pointer to service
	return &service{store}
}

func (s *service) CreateOrder(context.Context) error {
	return nil
}
