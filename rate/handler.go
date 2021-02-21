package rate

import (
	"log"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) GetRate(ctx context.Context, in *Rate) (*Response, error) {
	log.Printf("Receive message %s", in.source)
	return &Response{Rate: "bar"}, nil
}
