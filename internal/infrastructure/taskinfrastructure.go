package infrastructure

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap"
)

func TaskFindByClassifyId(ctx *contextx.Contextx, classifyId string) (result []model.TaskModel, err error) {
	result, err = ctx.BaseTx.TaskModel.FindByClassifyId(classifyId)
	return
}

func InsertTask(ctx *contextx.Contextx, tm *model.TaskModel) (id uint, err error) {
	id, err = ctx.BaseTx.TaskModel.Insert(tm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = i18n.NewCodeError(ctx.Language, module.TaskInsertErr, err.Error())
	}
	return
}
