package infrastructure

import (
	"errors"

	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TaskFindByClassifyId(ctx *contextx.Contextx, classifyId string) (result []entity.TaskEntity, err error) {
	result, err = ctx.BaseTx.TaskModel.FindByClassifyId(ctx.UID, classifyId)
	return
}

func TaskFirst(ctx *contextx.Contextx, p model.TaskModel) (result model.TaskModel, err error) {
	result, err = ctx.BaseTx.TaskModel.First(p)
	if err == nil {
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(module.TaskNotfoundErr)
		return
	}
	err = i18n.NewCodeError(module.TaskQueryErr)
	return
}

func TaskInsert(ctx *contextx.Contextx, tm *model.TaskModel) (oc string, err error) {
	// _, err = ClassifyFirst(ctx, model.ClassifyModel{OnlyCode: tm.ClassifyId})
	// if err != nil {
	// 	return
	// }

	_, err = TaskModeFirst(ctx, model.TaskModeModel{OnlyCode: tm.TaskModeId})
	if err != nil {
		return
	}

	tm.OnlyCode = utils.NewUlid()
	oc = tm.OnlyCode
	_, err = ctx.BaseTx.TaskModel.Insert(tm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = i18n.NewCodeError(module.TaskInsertErr)
	}
	return
}

func TaskUpdate(ctx *contextx.Contextx, tm model.TaskModel) (err error) {
	// _, err = ClassifyFirst(ctx, model.ClassifyModel{OnlyCode: tm.ClassifyId})
	// if err != nil {
	// 	return
	// }

	_, err = TaskModeFirst(ctx, model.TaskModeModel{OnlyCode: tm.TaskModeId})
	if err != nil {
		return
	}

	err = ctx.BaseTx.TaskModel.Update(&tm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = i18n.NewCodeError(module.TaskUpdateErr)
	}
	return
}

func TaskDelete(ctx *contextx.Contextx, id string) (err error) {
	err = ctx.BaseTx.TaskModel.Delete(id)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("oc", id))
		err = i18n.NewCodeError(module.TaskDeleteErr)
	}
	return
}
