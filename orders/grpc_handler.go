package main

import (
	"context"
	"log"

	pb "github.com/sikozonpc/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrderService
}

func NewGRPCHandler(grpcServer *grpc.Server, service OrderService) { // constructor
	// get from main.go(orders)
	handler := &grpcHandler{service: service}
	pb.RegisterOrderServiceServer(grpcServer, handler) // pass the handler to the client
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	// return nil, fmt.Errorf("some errors")
	log.Printf("New order received! Order ID: %v", p)
	o := &pb.Order{
		ID: "42",
	}

	return o, nil
}
