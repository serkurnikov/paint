package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "paint/internal/gPRC/imageProcessingService"
)

const (
	port = ":50051"
)

type server struct {
	pb.ImageProcessingServiceServer
}

func (s *server) MeanShiftFilter(ctx context.Context, in *pb.PyrMeanShiftFilteringParams) (*pb.PyrMeanShiftFilteringResponse, error) {
	log.Printf("Received: %v", in.In)
	return &pb.PyrMeanShiftFilteringResponse{
		Out:   "answer",
		Error: nil,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterImageProcessingServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
