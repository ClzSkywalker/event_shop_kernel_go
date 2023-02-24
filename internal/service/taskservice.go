package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap"
)

func InsertTask(m model.ITaskModel, tm *model.TaskModel) (id uint, err error) {
	id, err = m.Insert(tm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = i18n.NewCodeError(module.TaskInsertErr, err.Error())
	}
	return
}
