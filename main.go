package main

import (
	"log"
	"net"

	"github.com/mannamcc/grpc-server/rate"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := rate.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	rate.RegisterRateServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Printf("GRPC server is running")
}
