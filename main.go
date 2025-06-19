package main

import (
	_ "GFShop/internal/packed"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"

	_ "GFShop/internal/logic"

	"GFShop/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
