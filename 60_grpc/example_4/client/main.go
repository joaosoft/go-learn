package main

import (
	"log"

	pb "github.com/joaosoft/golang-learn/60_grpc/example_4"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//defer conn.Close()
	c := pb.NewServiceClient(conn)

	// Contact the server and print out its response.
	response, err := c.GetName(context.Background(), &pb.Request{Id: 123})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", response.Name)
}
