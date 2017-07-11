package grpc_sample

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// SampleRPC controller.
type SampleRPC struct{}

// NewSampleRPC creates a new gRPC controller and registers it on the provided gRPC API service.
func NewSampleRPC(server *grpc.Server) *SampleRPC {
	controller := &SampleRPC{}

	RegisterSampleServer(server, controller)

	return controller
}

// Ping ...
func (rpc *SampleRPC) Ping(ctx context.Context, req *EmptyRequest) (*PingResponse, error) {
	res := &PingResponse{
		Pong: "pong",
	}

	return res, nil
}
