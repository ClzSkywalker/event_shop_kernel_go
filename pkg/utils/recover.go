package utils

import (
	"runtime/debug"

	"go.uber.org/zap"
)

type recoverFunc func()

/**
 * @Author         : Angular
 * @Date           : 2023-02-06
 * @Description    : panic recover
 * @param           {recoverFunc} goFunc
 * @param           {logErrFunc} logFunc
 * @return          {*}
 */
func RecoverFunc(callBack recoverFunc) {
	defer func() {
		if err := recover(); err != nil {
			ZapLog.Panic("[server panic]", zap.Any("RecoverFunc", err))
			debug.PrintStack()
		}
	}()
	callBack()
}
