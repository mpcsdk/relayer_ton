// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "relayer_ton/api/relayer/v1"
)

type (
	IRelayer interface {
		Deploy(ctx context.Context, req *v1.DeployReq) (*v1.DeployRes, error)
	}
)

var (
	localRelayer IRelayer
)

func Relayer() IRelayer {
	if localRelayer == nil {
		panic("implement not found for interface IRelayer, forgot register?")
	}
	return localRelayer
}

func RegisterRelayer(i IRelayer) {
	localRelayer = i
}
