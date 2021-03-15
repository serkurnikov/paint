package main

import (
	"context"
	"gocv.io/x/gocv"
	"google.golang.org/grpc"
	"log"
	"net"
	"paint/internal/gRPC/imageProcessingService/colorProcessing/mixcolors"
	"paint/internal/gRPC/imageProcessingService/colorProcessing/prominentcolor"
	"paint/internal/gRPC/imageProcessingService/contoursProcessing"
	"paint/internal/gRPC/imageProcessingService/imageFilterProcessing"
	"paint/internal/gRPC/imageProcessingService/morphProcessing"
	pb "paint/internal/gRPC/imageProcessingService/service"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedImageProcessingServiceServer
}

func (s *server) PyrMeanShiftFiltering(ctx context.Context, in *pb.PyrRequest) (*pb.DefaultReply, error) {
	imageFilterProcessing.PyrMeanShiftFiltering(in.PathPicture, in.Sp, in.Sr, in.MaxLevel)
	return &pb.DefaultReply{OutPicture: ""}, nil
}

func (s *server) DrawDefaultContours(ctx context.Context, in *pb.ContoursRequest) (*pb.DefaultReply, error) {
	contoursProcessing.DrawDefaultContours(in.PathPicture, in.T1, in.T2)
	return &pb.DefaultReply{OutPicture: ""}, nil
}

func (s *server) DrawCustomContours(ctx context.Context, in *pb.ContoursRequest) (*pb.DefaultReply, error) {
	contoursProcessing.DrawCustomContours(in.PathPicture, in.T1, in.T2)
	return &pb.DefaultReply{OutPicture: ""}, nil
}

func (s *server) DrawHoughLinesWithParams(ctx context.Context, in *pb.HoughLinesWithPRequest) (*pb.DefaultReply, error) {
	contoursProcessing.DrawHoughLinesWithParams(in.PathPicture, in.Rho, in.Theta, int(in.Threshold), in.MinLineLength, in.MaxLineGap)
	return &pb.DefaultReply{OutPicture: ""}, nil
}

func (s *server) DrawHoughCircles(ctx context.Context, in *pb.HoughCirclesRequest) (*pb.DefaultReply, error) {
	contoursProcessing.DrawHoughCircles(in.PathPicture, gocv.HoughGradient, float64(in.Dp),
		float64(in.MinDist), float64(in.P1), float64(in.P2), int(in.MinR), int(in.MaxR))

	return &pb.DefaultReply{OutPicture: ""}, nil
}

func (s *server) Threshold(ctx context.Context, in *pb.ThresholdRequest) (*pb.DefaultReply, error) {
	morphProcessing.Threshold(in.PathPicture, in.Thresh, in.Maxvalue)

	return &pb.DefaultReply{OutPicture: ""}, nil
}

func (s *server) Watershed(ctx context.Context, in *pb.WatershedRequest) (*pb.DefaultReply, error) {
	morphProcessing.Watershed(in.PathPicture, int(in.NErode), int(in.NDilate))

	return &pb.DefaultReply{OutPicture: ""}, nil
}

func (s *server) Open(ctx context.Context, in *pb.OpenRequest) (*pb.DefaultReply, error) {
	morphProcessing.Open(in.PathPicture, int(in.KernelSize))

	return &pb.DefaultReply{OutPicture: ""}, nil
}

func (s *server) Close(ctx context.Context, in *pb.CloseRequest) (*pb.DefaultReply, error) {
	morphProcessing.Close(in.PathPicture, int(in.KernelSize))

	return &pb.DefaultReply{OutPicture: ""}, nil
}

func (s *server) FindBlendStructure(ctx context.Context, in *pb.SeparationRequest) (*pb.SeparationReply, error) {
	result := mixcolors.BlendStructureAmongFabricColors(in.MainColorS, in.ColorFabric, int(in.SpaceType))
	blendStructures := make([]*pb.SeparationReply_BlendStructure, 0)

	for i := 0; i < len(result); i++ {
		blendStructures = append(blendStructures,
		&pb.SeparationReply_BlendStructure{
			C1Hex:     result[i].C1Hex,
			C2Hex:     result[i].C2Hex,
			C3Hex:     result[i].C3Hex,
			C2Portion: result[i].C2Portion,
			C3Portion: result[i].C3Portion,
			ResultHex: result[i].ResultHex,
		})
	}
	return &pb.SeparationReply{
		BlendStructures: blendStructures,
	}, nil
}

func (s *server) DisplayPictureInDominatedColors(ctx context.Context, in *pb.PictureInDominatedColorsRequest) (*pb.DefaultReply, error) {
	prominentcolor.DisplayPictureInDominatedColors(in.InPicture, in.OutPicture, int(in.NumberOfClusters))

	return &pb.DefaultReply{OutPicture: ""}, nil
}

func (s *server) ColorSeparation(ctx context.Context, in *pb.SeparationRequest) (*pb.SeparationReply, error) {
	mixcolors.ColorSeparation(in.MainColorS, in.ColorFabric, int(in.SpaceType))

	return &pb.SeparationReply{BlendStructures: nil}, nil
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
