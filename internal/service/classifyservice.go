package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func QueryClassifyByUidAndTitle(ctx *contextx.Contextx, tx model.IClassifyModel,
	title string) (result model.ClassifyModel, err error) {
	result, err = tx.QueryByTitle(ctx.TID, ctx.UID, title)
	if err != nil && err != gorm.ErrRecordNotFound {
		loggerx.ZapLog.Error(err.Error(), zap.Any("ctx", ctx), zap.String("title", title))
		err = i18n.NewCodeError(ctx.Language, module.ClassifyQueryErr)
		return
	}
	return
}

func QueryAllClassify(tx model.IClassifyModel, t entity.TokenInfo) (cms []model.ClassifyModel, err error) {
	cms, err = tx.QueryAll(t.TID)
	return
}

func InsertClassify(ctx *contextx.Contextx, tx model.IClassifyModel, cm *model.ClassifyModel) (id uint, err error) {
	_, err = QueryClassifyByUidAndTitle(ctx, tx, cm.Title)
	if err != gorm.ErrRecordNotFound {
		err = i18n.NewCodeError(ctx.Language, module.ClassifyExistedErr)
		return
	}

	id, err = tx.Insert(cm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
		err = i18n.NewCodeError(ctx.Language, module.ClassifyInsertErr)
	}
	return
}

func UpdateClassify(ctx *contextx.Contextx, tx model.IClassifyModel, t entity.TokenInfo, cm model.ClassifyModel) (err error) {
	_, err = QueryClassifyByUidAndTitle(ctx, tx, cm.Title)
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

func DeleteClassify(ctx *contextx.Contextx, tx model.IClassifyModel, oc, uid string) (err error) {
	err = tx.Delete(oc, uid)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("oc", oc), zap.String("uid", uid))
		err = i18n.NewCodeError(ctx.Language, module.ClassifyUpdateErr)
	}
	return
}
