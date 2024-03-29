package infrastructure

import (
	"errors"

	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TaskFirst(ctx *contextx.Contextx, p model.TaskModel) (result model.TaskModel, err error) {
	result, err = ctx.BaseTx.TaskModel.First(p)
	if err == nil {
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errorx.NewCodeError(module.TaskNotfoundErr)
		return
	}
	err = errorx.NewCodeError(module.TaskQueryErr)
	return
}

func TaskInsert(ctx *contextx.Contextx, tm *model.TaskModel) (oc string, err error) {
	_, err = DevideFirst(ctx, model.DevideModel{OnlyCode: tm.DevideId})
	if err != nil {
		return
	}

	_, err = TaskModeFirst(ctx, model.TaskModeModel{OnlyCode: tm.TaskModeId})
	if err != nil {
		return
	}

	tm.OnlyCode = utils.NewUlid()
	oc = tm.OnlyCode
	_, err = ctx.BaseTx.TaskModel.Insert(tm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = errorx.NewCodeError(module.TaskInsertErr)
	}
	return
}

func TaskUpdate(ctx *contextx.Contextx, tm model.TaskModel) (err error) {
	_, err = DevideFirst(ctx, model.DevideModel{OnlyCode: tm.DevideId})
	if err != nil {
		return
	}

	_, err = TaskModeFirst(ctx, model.TaskModeModel{OnlyCode: tm.TaskModeId})
	if err != nil {
		return
	}

	err = ctx.BaseTx.TaskModel.Update(&tm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = errorx.NewCodeError(module.TaskUpdateErr)
	}
	return
}

func TaskDelete(ctx *contextx.Contextx, id string) (err error) {
	err = ctx.BaseTx.TaskModel.Delete(id)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("oc", id))
		err = errorx.NewCodeError(module.TaskDeleteErr)
	}
	return
}
