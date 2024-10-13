package main

import (
	"context"
	"log"
	"net"

	common "github.com/sikozonpc/commons"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer( // gRPC has default limits on the size of frames and messages
		grpc.MaxRecvMsgSize(1024*1024*50), // 50 MB
		grpc.MaxSendMsgSize(1024*1024*50), // 50 MB
	) // *grpc.Server

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer l.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer, svc)

	svc.CreateOrder(context.Background())
	log.Println("GRPC Server Started at", grpcAddr)

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}

}
