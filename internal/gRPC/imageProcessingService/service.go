package imageProcessingService

import (
	"context"
	"log"
	"net"
	"paint/internal/gRPC/imageProcessingService/imageFilterProcessing"

	"google.golang.org/grpc"
	pb "paint/internal/gRPC/imageProcessingService/service"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedImageProcessingServiceServer
}

func (s *server) PyrMeanShiftFiltering(ctx context.Context, in *pb.PyrMeanShiftFilteringRequest) (*pb.PyrMeanShiftFilteringReply, error) {
	log.Printf("Received: %v", in.In)
	imageFilterProcessing.PyrMeanShiftFiltering(in.In, "", []float64{0,0})
	return &pb.PyrMeanShiftFilteringReply{Message: "Result = "}, nil
}

func RunImageProcessingServer() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return err
}
