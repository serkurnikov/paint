package paintApi

import (
	"context"
	"gocv.io/x/gocv"
	api "paint/pkg/api/proto_files"
	"paint/pkg/imageProcessingService/contoursProcessing"
)

func (s service) DrawDefaultContours(ctx context.Context, in *api.ContoursRequest) (*api.DefaultReply, error) {
	//contoursProcessing.DrawDefaultContours(in.PathPicture, in.T1, in.T2)
	return &api.DefaultReply{OutPicture: "TEST"}, nil
}

func (s service) DrawCustomContours(ctx context.Context, in *api.ContoursRequest) (*api.DefaultReply, error) {
	contoursProcessing.DrawCustomContours(in.PathPicture, in.T1, in.T2)
	return &api.DefaultReply{OutPicture: ""}, nil
}

func (s service) DrawHoughLinesWithParams(ctx context.Context, in *api.HoughLinesWithPRequest) (*api.DefaultReply, error) {
	contoursProcessing.DrawHoughLinesWithParams(in.PathPicture, in.Rho, in.Theta, int(in.Threshold), in.MinLineLength, in.MaxLineGap)
	return &api.DefaultReply{OutPicture: ""}, nil
}

func (s service) DrawHoughCircles(ctx context.Context, in *api.HoughCirclesRequest) (*api.DefaultReply, error) {
	contoursProcessing.DrawHoughCircles(in.PathPicture, gocv.HoughGradient, float64(in.Dp),
		float64(in.MinDist), float64(in.P1), float64(in.P2), int(in.MinR), int(in.MaxR))

	return &api.DefaultReply{OutPicture: ""}, nil
}
