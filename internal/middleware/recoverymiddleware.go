package middleware

import (
	"runtime"

	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryMiddleware(c *gin.Context, err interface{}) {
	pc, file, line, ok := runtime.Caller(3)
	funcName := runtime.FuncForPC(pc).Name()
	if !ok {
		loggerx.ZapLog.Panic("panic", zap.Any("err", err))
		return
	}
	loggerx.ZapLog.Panic("panic", zap.String("file", file), zap.Int("line", line), zap.String("func", funcName), zap.Any("err", err))
}
