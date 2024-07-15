// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package relayer

import (
	"context"

	"relayer_ton/api/relayer/v1"
)

type IRelayerV1 interface {
	Deploy(ctx context.Context, req *v1.DeployReq) (res *v1.DeployRes, err error)
}
