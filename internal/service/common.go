package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
)

func StructToStruct(ctx *contextx.Contextx, t, v interface{}) (err error) {
	err = utils.StructToStruct(t, v)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(module.StructToStructErr)
		return
	}
	return
}
