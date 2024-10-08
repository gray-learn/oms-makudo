package main

import (
	"log"
	"net/http"

	common "github.com/sikozonpc/commons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// httpAddr = ":8080" // const
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = "localhost:3000"
)

func main() {

	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial order service: %v", err)
	}
	defer conn.Close()

	log.Println("Dialing orders service at", orderServiceAddr)

	mux := http.NewServeMux() // http router
	handle := NewHandler(c)
	handle.registerRoutes(mux)

	log.Printf("start http server at %s", httpAddr)

	// start http server
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("falied to start http server", err)
	}
}
