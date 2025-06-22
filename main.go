package main

import (
	_ "GFShop/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	"GFShop/internal/cmd"
	_ "GFShop/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
