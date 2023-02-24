package processx

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap/zapcore"
)

func HandleSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	s := <-c

	loggerx.ZapLog.Info("received os signal, exit kernel process now", zapcore.Field{Key: "signal", Interface: s})
}
