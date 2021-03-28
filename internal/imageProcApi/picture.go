package imageProcApi

import (
	"context"
	"log"
	api "paint/api/proto/pb"
)

func (i imageProcApi) PyrMeanShiftFilter(ctx context.Context, in *api.PyrRequest) (*api.DefaultReply, error) {
	req := api.PyrRequest{
		PathPicture: in.PathPicture,
		Sp:          in.Sp,
		Sr:          in.Sr,
		MaxLevel:    in.MaxLevel,
	}
	_, err := i.processingServiceClient.PyrMeanShiftFiltering(ctx, &req)
	if err != nil {
		log.Printf("error result pyrMeanShiftFilter")
	}
	return &api.DefaultReply{OutPicture: "TEST"}, nil
}