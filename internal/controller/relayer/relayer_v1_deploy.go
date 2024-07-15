package relayer

import (
	"context"

	v1 "relayer_ton/api/relayer/v1"
	"relayer_ton/internal/service"
)

func (c *ControllerV1) Deploy(ctx context.Context, req *v1.DeployReq) (*v1.DeployRes, error) {

	res, err := service.Relayer().Deploy(ctx, req)
	return res, err
}
