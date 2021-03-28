package app

import (
	"context"
	api "paint/api/proto/pb"
)

func (a App) PyrMeanShiftFilter(ctx context.Context, in *api.PyrRequest) (*api.DefaultReply, error) {
	return a.imageProcApi.PyrMeanShiftFilter(ctx, in)
}