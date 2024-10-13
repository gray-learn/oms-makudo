package main

import (
	"context"
	pb "github.com/sikozonpc/commons/api"
)

type OrderService interface {
	CreateOrder(context.Context) error // payload and output
	validateOrder(context.Context, *pb.CreateOrderRequest) error
}

type OrderStore interface {
	Create(context.Context) error // payload and output
}
