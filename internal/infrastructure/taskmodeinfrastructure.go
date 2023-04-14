package infrastructure

import (
	"errors"

	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func TaskModeFirst(ctx *contextx.Contextx, tm model.TaskModeModel) (result model.TaskModeModel, err error) {
	result, err = ctx.BaseTx.TaskModeModel.First(tm)
	if err == nil {
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errorx.NewCodeError(module.TaskModeNotFoundErr)
		return
	}
	err = errorx.NewCodeError(module.TaskModeQueryErr)
	loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
	return
}

func InsertTaskMode(ctx *contextx.Contextx, tm *model.TaskModeModel) (id uint, err error) {
	id, err = ctx.BaseTx.TaskModeModel.Insert(tm)
	return
}
