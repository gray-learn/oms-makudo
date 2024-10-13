package main

import (
	"log"
	"net/http"

	common "github.com/sikozonpc/commons"
	pb "github.com/sikozonpc/commons/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// httpAddr = ":8080" // const
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = "localhost:3000"
)

func main() {

	// proxyCA := "/var/tmp/fullchain.pem" // CA cert that signed the proxy
	// f, err := os.ReadFile(proxyCA)
	// p := x509.NewCertPool()
	// p.AppendCertsFromPEM(f)
	// tlsConfig := &tls.Config{
	// 	RootCAs: p,
	// }

	conn, err := grpc.NewClient(orderServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(1024*1024*50), // 50 MB
		),
		// orderServiceAddr,
		// grpc.WithTransportCredentials(insecure.NewCredentials())
	)

	if err != nil {
		log.Fatalf("failed to dial order service: %v", err)
	}
	defer conn.Close()

	log.Println("Dialing orders service at", orderServiceAddr)
	c := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux() // http router
	handle := NewHandler(c)
	handle.registerRoutes(mux)

	log.Printf("start http server at %s", httpAddr)

	// start http server
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("falied to start http server", err)
	}
}
