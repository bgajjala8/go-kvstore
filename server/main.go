package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/bgajjala8/go-cache/internal/service"
	pb "github.com/bgajjala8/go-cache/proto-public/go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr     = "localhost:50051"
	httpAddr = "localhost:8080"
)

func main() {
	// Start gRPC server in a separate goroutine
	go func() {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("Failed to listen: %v\n", err)
		}
		log.Printf("gRPC server listening at %v\n", lis.Addr())

		s := grpc.NewServer()
		svc := service.NewService()
		pb.RegisterCacheServiceServer(s, svc)

		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v\n", err)
		}
	}()

	// Create a new gRPC Gateway mux for the HTTP server on main goroutine
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterCacheServiceHandlerFromEndpoint(context.Background(), mux, addr, opts)
	if err != nil {
		log.Fatalf("Failed to start HTTP gateway: %v", err)
	}

	// Start the HTTP server
	log.Printf("Starting HTTP server on %s\n", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("Failed to serve HTTP: %v", err)
	}
}
