package paintApi

import (
	"context"
	api "paint/pkg/api/proto_files"
	"paint/pkg/imageProcessingService/morphProcessing"
)

func (s service) Threshold(ctx context.Context, in *api.ThresholdRequest) (*api.DefaultReply, error) {
	morphProcessing.Threshold(in.PathPicture, in.Thresh, in.Maxvalue)

	return &api.DefaultReply{OutPicture: ""}, nil
}

func (s service) Watershed(ctx context.Context, in *api.WatershedRequest) (*api.DefaultReply, error) {
	morphProcessing.Watershed(in.PathPicture, int(in.NErode), int(in.NDilate))

	return &api.DefaultReply{OutPicture: ""}, nil
}

func (s service) Open(ctx context.Context, in *api.OpenRequest) (*api.DefaultReply, error) {
	morphProcessing.Open(in.PathPicture, int(in.KernelSize))

	return &api.DefaultReply{OutPicture: ""}, nil
}

func (s service) Close(ctx context.Context, in *api.CloseRequest) (*api.DefaultReply, error) {
	morphProcessing.Close(in.PathPicture, int(in.KernelSize))

	return &api.DefaultReply{OutPicture: ""}, nil
}
