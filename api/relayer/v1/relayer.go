package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeployReq struct {
	g.Meta  `path:"/deploy" tags:"Deploy" method:"get" summary:"You first Deploy api"`
	Address string `json:"address"`
}
type DeployRes struct {
	g.Meta  `mime:"text/html" example:"string"`
	Address string `json:"address"`
}
