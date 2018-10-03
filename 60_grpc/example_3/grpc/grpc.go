package grpc

import (
	"net"

	"fmt"
	"log"

	"google.golang.org/grpc"
)

// GRPC controller.
type GRPC struct {
	Server  *grpc.Server
	address string
	running bool
}

// NewGRPC creates a new gRPC controller and respective API.
// TODO: convert parameters to a Configuration struct
func NewGRPC(address string) *GRPC {
	opts := make([]grpc.ServerOption, 1)
	opts[0] = grpc.MaxConcurrentStreams(10)

	// TODO: add support for TLS
	// creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// if err != nil {
	// 	grpclog.Fatalf("Failed to generate credentials %v", err)
	// }
	// opts[1] = grpc.Creds(creds)

	controller := &GRPC{
		Server:  grpc.NewServer(opts...),
		address: address,
		running: false,
	}

	return controller
}

// Start initializes the gRPC server. This is a blocking operation.
func (rpc *GRPC) Start() error {
	bind, err := net.Listen("tcp", rpc.address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Fatalf("starting gRPC server on [addr:%s]", rpc.address)

	rpc.running = true
	if err := rpc.Server.Serve(bind); err != nil {
		rpc.running = false
		return err
	}
	rpc.running = false
	log.Fatalf("stopped gRPC server on [addr:%s]", rpc.address)

	return nil
}

// Stop attempts to shutdown gracefully.
func (rpc *GRPC) Stop() error {
	if rpc.running != true {
		return fmt.Errorf("rpc server is already stopped")
	}

	rpc.Server.GracefulStop()
	rpc.running = false

	return nil
}

// IsRunning returns true if the gRPC service is running and listening for connections.
func (rpc *GRPC) IsRunning() bool {
	return rpc.running
}
