package utils

import (
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap/zapcore"
)

func HandleSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	s := <-c

	ZapLog.Info("received os signal, exit kernel process now", zapcore.Field{Key: "signal", Interface: s})
}
