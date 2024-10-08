package main

import "context"

type OrderService interface {
	CreateOrder(context.Context) error // payload and output
}

type OrderStore interface {
	Create(context.Context) error // payload and output
}
