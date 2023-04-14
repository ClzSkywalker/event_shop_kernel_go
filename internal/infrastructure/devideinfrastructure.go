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
	err = ctx.BaseTx.DevideModel.Update(m)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = errorx.NewCodeError(module.DevideUpdateErr)
		return
	}
	return
}

func DevideDelete(ctx *contextx.Contextx, oc string) (err error) {
	err = ctx.BaseTx.DevideModel.Delete(oc)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = errorx.NewCodeError(module.DevideDeleteErr)
	}
	return
}
