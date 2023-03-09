package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/contextx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap"
)

func InsertTask(ctx *contextx.Contextx, m model.ITaskModel, tm *model.TaskModel) (id uint, err error) {
	id, err = m.Insert(tm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = i18n.NewCodeError(ctx.Language, module.TaskInsertErr, err.Error())
	}
	return
}
