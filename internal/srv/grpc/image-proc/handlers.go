package image_proc

import (
	"context"
	api "paint/api/proto/pb"
	"paint/pkg/imageProcessingService/imageFilterProcessing"
)

func (s *service) PyrMeanShiftFiltering(ctx context.Context, in *api.PyrRequest) (*api.DefaultReply, error) {
	imageFilterProcessing.PyrMeanShiftFiltering(in.PathPicture, in.Sp, in.Sr, in.MaxLevel)
	return &api.DefaultReply{OutPicture: ""}, nil
}
