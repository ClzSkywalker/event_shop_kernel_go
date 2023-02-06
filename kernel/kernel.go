package kernel

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/server"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	_ "golang.org/x/mobile/bind"
)

func StartKernel() {
	go utils.RecoverFunc(server.ServerStart)
}
