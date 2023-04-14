package infrastructure

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
)

func TaskContentInsert(ctx *contextx.Contextx, content string) (oc string, err error) {
	oc = utils.NewUlid()
	m := model.TaskContentModel{
		OnlyCode: oc,
		Content:  content,
	}
	_, err = ctx.BaseTx.TaskContentModel.Insert(m)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = errorx.NewCodeError(module.TaskContentInsertErr)
		return
	}
	return
}
