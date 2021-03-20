package paintApi

import (
	"context"
	api "paint/pkg/api/proto_files"
	"paint/pkg/imageProcessingService/colorProcessing/mixcolors"
)

func (s service) FindBlendStructure(ctx context.Context, in *api.SeparationRequest) (*api.SeparationReply, error) {
	result := mixcolors.BlendStructureAmongFabricColors(in.MainColorS, in.ColorFabric, int(in.SpaceType))
	blendStructures := make([]*api.SeparationReply_BlendStructure, 0)

	for i := 0; i < len(result); i++ {
		blendStructures = append(blendStructures,
			&api.SeparationReply_BlendStructure{
				C1Hex:     result[i].C1Hex,
				C2Hex:     result[i].C2Hex,
				C3Hex:     result[i].C3Hex,
				C2Portion: result[i].C2Portion,
				C3Portion: result[i].C3Portion,
				ResultHex: result[i].ResultHex,
			})
	}
	return &api.SeparationReply{
		BlendStructures: blendStructures,
	}, nil
}

func (s service) ColorSeparation(ctx context.Context, in *api.SeparationRequest) (*api.SeparationReply, error) {
	mixcolors.ColorSeparation(in.MainColorS, in.ColorFabric, int(in.SpaceType))

	return &api.SeparationReply{BlendStructures: nil}, nil
}
