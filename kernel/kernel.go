package kernel

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/server"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	_ "golang.org/x/mobile/bind"
)

func StartKernel(port, local int, mode, dbPath string) {
	go utils.RecoverFunc(func() { server.KernelServer(port, local, mode, dbPath) })
}
