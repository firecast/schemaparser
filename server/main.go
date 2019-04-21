package main

import (
	"context"
	"log"
	"net"

	pb "github.com/firecast/schemaparser/protobufs"
	"github.com/firecast/schemaparser/protobufs/types/thing"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) GetSchemaForURL(ctx context.Context, in *pb.ParserRequest) (*thing.Product, error) {
	log.Printf("Received: %v", in.Url)
	return &thing.Product{Identifier: "1", Url: in.Url}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSchemaParserServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
