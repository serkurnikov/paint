package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"paint/internal/app"
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
	imageFilterProcessing.PyrMeanShiftFiltering(in.PathPicture, in.Sp, in.Sr, in.MaxLevel)
	return &pb.PyrReply{OutPicture: ""}, nil
}

func (s *server) DrawCountours(ctx context.Context, in *pb.CountourRequest) (*pb.CountourReply, error) {
	app.FindingMatchingGeometricShapes(in.PathPicture)
	return &pb.CountourReply{OutPicture: ""}, nil
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
