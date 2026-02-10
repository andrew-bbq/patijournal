package main

import (
	"log"
	"net"
	"patijournal/internal/repository"
	"patijournal/internal/service"
	"patijournal/pkg/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	connString := "host=localhost port=5432 user=postgres password=password dbname=postgres sslmode=disable"
	repo, err := repository.NewEntryPostgresRepository(connString)
	if err != nil {
		log.Fatalf("Failed to create repository: %v", err)
	}
	entryService := service.NewEntryService(repo)
	pb.RegisterEntryServiceServer(grpcServer, entryService)

	reflection.Register(grpcServer)

	log.Printf("gRPC server listening on %s", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
