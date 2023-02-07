package mobile

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/server"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	_ "golang.org/x/mobile/bind"
)

func StartKernel(port, local int, mode, dbPath, logPath string) {
	go utils.RecoverFunc(func() {
		server.KernelServer(container.AppConfig{
			Port:     port,
			Mode:     mode,
			DbPath:   dbPath,
			Language: local,
			LogPath:  logPath,
		})
	})
}
