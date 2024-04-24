package main

import (
	"log"
	"net"

	pb "github.com/joaosoft/golang-learn/60_grpc/example_4"

	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port         = ":50051"
	default_name = "joao ribeiro"
)

type server struct{}

// GetName implements service.GetName
func (s *server) GetName(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Name: default_name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("LOADED SERVER")
}
