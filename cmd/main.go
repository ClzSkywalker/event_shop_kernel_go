package main

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/server"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
)

func main() {
	go server.Serve()
	utils.HandleSignal()
}
