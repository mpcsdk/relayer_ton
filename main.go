package main

import (
	_ "relayer_ton/internal/packed"

	_ "relayer_ton/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"relayer_ton/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
