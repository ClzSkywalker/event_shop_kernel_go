package mobile

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/server"
	"github.com/clz.skywalker/event.shop/kernal/pkg/recoverx"
	_ "golang.org/x/mobile/bind"
)

func StartKernel(port int, mode, dbPath, logPath string) {
	go recoverx.RecoverFunc(func() {
		server.KernelServer(container.AppConfig{
			Port:    port,
			Mode:    mode,
			DbPath:  dbPath,
			LogPath: logPath,
		})
	})
}
