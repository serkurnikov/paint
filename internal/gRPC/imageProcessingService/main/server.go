package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"paint/internal/gRPC/imageProcessingService/imageFilterProcessing"
	pb "paint/internal/gRPC/imageProcessingService/service"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedImageProcessingServiceServer
}

func (s *server) PyrMeanShiftFiltering(ctx context.Context, in *pb.PyrRequest) (*pb.PyrReply, error) {
	log.Printf("PyrMeanShiftFiltering: %v", in.PathPicture)
	result := imageFilterProcessing.PyrMeanShiftFiltering(in.PathPicture, in.Sp, in.Sr, in.MaxLevel)
	return &pb.PyrReply{OutPicture: "Result placed in " + result}, nil
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
