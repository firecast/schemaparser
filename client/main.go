package main

import (
	"context"
	"log"
	"time"

	pb "github.com/firecast/schemaparser/protobufs"
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
	defer conn.Close()
	c := pb.NewSchemaParserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetSchemaForURL(ctx, &pb.ParserRequest{Url: "https://www.google.com"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("%s", r)
}
