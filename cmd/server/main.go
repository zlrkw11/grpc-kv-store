package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/rayzhao/grpc-kv-store/internal/server"
	"github.com/rayzhao/grpc-kv-store/internal/store"
	kvstorev1 "github.com/rayzhao/grpc-kv-store/pkg/kvstore/v1"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	st := store.New()
	srv := server.New(st)
	kvstorev1.RegisterKVStoreServer(grpcServer, srv)

	log.Println("gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
