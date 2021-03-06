package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "paint/internal/gPRC/imageProcessingService"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewImageProcessingServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.PyrMeanShiftFilter(ctx, &pb.PyrMeanShiftFilteringParams{
		In:       "request",
	})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Out image: %s", r.Out)
}
