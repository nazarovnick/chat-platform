package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nazarovnick/chat-platform/services/user/internal/controller/grpc/v1/user"
	"github.com/nazarovnick/chat-platform/services/user/internal/controller/http"
	pb "github.com/nazarovnick/chat-platform/services/user/pkg/api/grpc/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"sync"
)

// main initializes the HTTP router and starts the server on specified port
func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	impl, err := user.NewServer()
	if err != nil {
		log.Fatalf("failed to create impl: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcServer := grpc.NewServer()
		pb.RegisterUserServiceServer(grpcServer, impl)
		reflection.Register(grpcServer)
		lis, err := net.Listen("tcp", ":8091")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC impl listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		gwMux := runtime.NewServeMux()

		if err = pb.RegisterUserServiceHandlerServer(ctx, gwMux, impl); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		httpServer := http.NewRouter(gwMux)
		addr := ":8080"
		log.Printf("HTTP server running at %s", addr)
		if err := httpServer.Listen(addr); err != nil {
			log.Fatalf("Failed to start impl: %v", err)
		}
	}()

	wg.Wait()
}
