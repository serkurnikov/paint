package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "paint/internal/gRPC/imageProcessingService/service"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedImageProcessingServiceServer
}

func (s *server) PyrMeanShiftFiltering(ctx context.Context, in *pb.PyrRequest) (*pb.PyrReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.PyrReply{Message: "Hello " + in.GetName()}, nil
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
