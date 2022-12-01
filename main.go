package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "mall-api/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	_ "mall-api/internal/logic"

	"mall-api/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
