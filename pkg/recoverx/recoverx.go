package recoverx

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/clz.skywalker/event.shop/kernal/pkg/consts"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap"
)

/**
 * @Author         : Angular
 * @Date           : 2023-02-06
 * @Description    : panic recover
 * @param           {recoverFunc} goFunc
 * @param           {logErrFunc} logFunc
 * @return          {*}
 */
func RecoverFunc(callBack consts.RecoverFunc) {
	defer func() {
		if err := recover(); err != nil {
			loggerx.ZapLog.Panic("[server panic]", zap.Any("RecoverFunc", err))
			debug.PrintStack()
		}
	}()
	callBack()
}

func RecoverReadChanFunc(logPath string, callBack consts.RecoverReadChanFunc, ch <-chan consts.DbInitStateType) {
	defer func() {
		if err := recover(); err != nil {
			buff := make([]byte, 1<<10)
			runtime.Stack(buff, false)
			_ = RewriteStderrFile(logPath, buff, err)
		}
	}()
	callBack(ch)
}

func RecoverWriteChanFunc(logPath string, callBack consts.RecoverWriteChanFunc, ch chan<- consts.DbInitStateType) {
	defer func() {
		if err := recover(); err != nil {
			buff := make([]byte, 1<<10)
			runtime.Stack(buff, false)
			_ = RewriteStderrFile(logPath, buff, err)
		}
	}()
	callBack(ch)
}

func RewriteStderrFile(path string, data []byte, msg interface{}) error {
	file, err := os.OpenFile(path+"/.panic.log", os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	_, err = file.WriteString(time.Now().Format(string(consts.DateTimeLayout)) + "\n" +
		"[panic msg]" + "\n" + fmt.Sprintf("%s\n", msg) +
		"[stack]" + "\n" + string(data) + "\n")
	if err != nil {
		return err
	}
	return nil
}
