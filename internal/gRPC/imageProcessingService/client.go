package imageProcessingService

import (
	"google.golang.org/grpc"
	"log"
	pb "paint/internal/gRPC/imageProcessingService/service"
)

const (
	address = "localhost:50051"
)

func NewImageProcessingClient() pb.ImageProcessingServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	return pb.NewImageProcessingServiceClient(conn)
}