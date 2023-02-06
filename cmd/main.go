package main

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/server"
	"github.com/clz.skywalker/event.shop/kernal/pkg/logger"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"go.uber.org/zap"
)

func main() {
	go utils.RecoverFunc(server.ServerStart, logPanic)
	utils.HandleSignal()
}

func logPanic(err interface{}) {
	logger.ZapLog.Panic("server panic", zap.Any("err", err))
}
