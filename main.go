package main

import (
	_ "mall-api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"mall-api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
