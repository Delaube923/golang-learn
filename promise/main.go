package main

import (
	_ "promise/internal/packed"

	_ "promise/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"promise/internal/cmd"
)

func main() {
	go cmd.WebSocketTask.Run(gctx.New())
	cmd.Main.Run(gctx.New())
}
