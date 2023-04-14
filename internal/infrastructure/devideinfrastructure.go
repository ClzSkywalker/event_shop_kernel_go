package infrastructure

import (
	"errors"

	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/gorm"
)

func DevideFirst(ctx *contextx.Contextx, where model.DevideModel) (result model.DevideModel, err error) {
	result, err = ctx.BaseTx.DevideModel.First(where)
	if err == nil {
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errorx.NewCodeError(module.DevideNotfoundErr)
		return
	}
	err = errorx.NewCodeError(module.DevideQueryErr)
	return
}

func DevideInsert(ctx *contextx.Contextx, m model.DevideModel) (oc string, err error) {
	_, err = DevideFirst(ctx, model.DevideModel{ClassifyId: m.ClassifyId, Title: m.Title})
	if err != nil && !errorx.Is(err, module.DevideNotfoundErr) {
		return
	} else if err == nil {
		err = errorx.NewCodeError(module.DevideTitleRepeatErr)
		return
	}

	_, err = ClassifyFirst(ctx, model.ClassifyModel{CreatedBy: ctx.UID, OnlyCode: m.ClassifyId})
	if err != nil {
		return
	}

	oc = utils.NewUlid()
	m.OnlyCode = oc
	m.CreatedBy = ctx.UID
	_, err = ctx.BaseTx.DevideModel.Insert(&m)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = errorx.NewCodeError(module.DevideInsertErr)
		return
	}
	return
}

func DevideUpdate(ctx *contextx.Contextx, m model.DevideModel) (err error) {
	de, err := DevideFirst(ctx, model.DevideModel{ClassifyId: m.ClassifyId, Title: m.Title})
	if err != nil && !errorx.Is(err, module.DevideNotfoundErr) {
		return
	} else if err == nil && de.OnlyCode != m.OnlyCode {
		err = errorx.NewCodeError(module.DevideTitleRepeatErr)
		return
	}

	_, err = ClassifyFirst(ctx, model.ClassifyModel{CreatedBy: ctx.UID, OnlyCode: m.ClassifyId})
	if err != nil {
		return
	}

	err = ctx.BaseTx.DevideModel.Update(m)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = errorx.NewCodeError(module.DevideUpdateErr)
		return
	}
	return
}

func DevideDelete(ctx *contextx.Contextx, oc string) (err error) {
	_, err = TaskFirst(ctx, model.TaskModel{DevideId: oc})
	if err != nil && !errorx.Is(err, module.TaskNotfoundErr) {
		return
	} else if err == nil {
		err = errorx.NewCodeError(module.DevideDelExistTaskErr)
		return
	}
	err = ctx.BaseTx.DevideModel.Delete(oc)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = errorx.NewCodeError(module.DevideDeleteErr)
	}
	return
}
