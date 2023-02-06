package kernel

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/server"
	_ "golang.org/x/mobile/bind"
)

func StartKernel() {
	go server.ServerStart()
}
