package paintApi

import (
	"context"
	api "paint/pkg/api/proto_files"
	"paint/pkg/imageProcessingService/colorProcessing/prominentcolor"
	"paint/pkg/imageProcessingService/imageFilterProcessing"
)

func (s service) PyrMeanShiftFiltering(ctx context.Context, in *api.PyrRequest) (*api.DefaultReply, error) {
	imageFilterProcessing.PyrMeanShiftFiltering(in.PathPicture, in.Sp, in.Sr, in.MaxLevel)
	return &api.DefaultReply{OutPicture: ""}, nil
}

func (s service) DisplayPictureInDominatedColors(ctx context.Context, in *api.PictureInDominatedColorsRequest) (*api.DefaultReply, error) {
	prominentcolor.DisplayPictureInDominatedColors(in.InPicture, in.OutPicture, int(in.NumberOfClusters))

	return &api.DefaultReply{OutPicture: ""}, nil
}
