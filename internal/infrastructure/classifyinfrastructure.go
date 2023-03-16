package infrastructure

import (
	"errors"

	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func ClassifyFindByTeamId(ctx *contextx.Contextx, tx model.IClassifyModel) (cms []model.ClassifyModel, err error) {
	cms, err = tx.FindByTeamId(ctx.TID)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("ctx", ctx))
		err = i18n.NewCodeError(ctx.Language, module.ClassifyQueryErr)
		return
	}
	return
}

func ClassifyInsert(ctx *contextx.Contextx, tx model.IClassifyModel, cm *model.ClassifyModel) (cid string, err error) {
	_, err = tx.First(model.ClassifyModel{TeamId: cm.TeamId})
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(ctx.Language, module.ClassifyExistedErr)
		return
	}

	oc := utils.NewUlid()
	cm.OnlyCode = oc
	_, err = tx.Insert(cm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
		err = i18n.NewCodeError(ctx.Language, module.ClassifyInsertErr)
	}
	cid = oc
	return
}

func ClassifyUpdate(ctx *contextx.Contextx, tx model.IClassifyModel, cm model.ClassifyModel) (err error) {
	_, err = tx.First(cm)
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.ClassifyNotfoundErr)
		return
	}
	err = tx.Update(cm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
		err = i18n.NewCodeError(ctx.Language, module.ClassifyUpdateErr)
	}
	return
}

func ClassifyDel(ctx *contextx.Contextx, tx model.IClassifyModel, oc string) (err error) {
	err = tx.Delete(oc, ctx.UID)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("oc", oc), zap.String("tid", ctx.TID))
		err = i18n.NewCodeError(ctx.Language, module.ClassifyUpdateErr)
	}
	return
}
