package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nazarovnick/chat-platform/services/chat/internal/controller/grpc/v1/chat"
	httphandler "github.com/nazarovnick/chat-platform/services/chat/internal/controller/http"
	pb "github.com/nazarovnick/chat-platform/services/chat/pkg/api/grpc/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"sync"
)

// main initializes the HTTP router and starts the server on specified port
func main() {
	//app := http.NewRouter()
	//addr := ":8080"
	//log.Printf("Server running at %s", addr)
	//if err := app.Listen(addr); err != nil {
	//	log.Fatalf("Failed to start server: %v", err)
	//}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	server, err := chat.NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcServer := grpc.NewServer()
		pb.RegisterChatServiceServer(grpcServer, server)
		reflection.Register(grpcServer)
		lis, err := net.Listen("tcp", ":8091")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("gRPC server listening at %v", lis.Addr())
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		gwMux := runtime.NewServeMux()

		if err = pb.RegisterChatServiceHandlerServer(ctx, gwMux, server); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

		router := httphandler.NewRouter(gwMux)
		addr := ":8092"
		log.Printf("HTTP server listening at %s", addr)
		if err := http.ListenAndServe(addr, router); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	wg.Wait()
}
