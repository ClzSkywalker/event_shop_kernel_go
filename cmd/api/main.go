package main

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/server"
	"github.com/clz.skywalker/event.shop/kernal/pkg/processx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/recoverx"
)

func main() {
	go recoverx.RecoverFunc(server.CmdServer)
	processx.HandleSignal()
}
