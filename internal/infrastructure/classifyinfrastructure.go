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

func ClassifyFindByTeamId(ctx *contextx.Contextx) (cms []model.ClassifyModel, err error) {
	cms, err = ctx.BaseTx.ClassifyModel.FindByTeamId(ctx.TID)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("ctx", ctx))
		err = i18n.NewCodeError(module.ClassifyQueryErr)
		return
	}
	return
}

func ClassifyFirst(ctx *contextx.Contextx, cm model.ClassifyModel) (result model.ClassifyModel, err error) {
	result, err = ctx.BaseTx.ClassifyModel.First(cm)
	if err == nil {
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(module.ClassifyNotfoundErr)
	}
	err = i18n.NewCodeError(module.ClassifyQueryErr)
	loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
	return
}

func ClassifyInsert(ctx *contextx.Contextx, cm *model.ClassifyModel) (cid string, err error) {
	_, err = ctx.BaseTx.ClassifyModel.First(model.ClassifyModel{TeamId: cm.TeamId})
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(module.ClassifyExistedErr)
		return
	}

	oc := utils.NewUlid()
	cm.OnlyCode = oc
	_, err = ctx.BaseTx.ClassifyModel.Insert(cm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
		err = i18n.NewCodeError(module.ClassifyInsertErr)
	}
	cid = oc
	return
}

func ClassifyUpdate(ctx *contextx.Contextx, cm model.ClassifyModel) (err error) {
	_, err = ctx.BaseTx.ClassifyModel.First(cm)
	if err != nil {
		err = i18n.NewCodeError(module.ClassifyNotfoundErr)
		return
	}
	err = ctx.BaseTx.ClassifyModel.Update(cm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
		err = i18n.NewCodeError(module.ClassifyUpdateErr)
	}
	return
}

//
// Author         : ClzSkywalker
// Date           : 2023-04-06
// Description    : 更新排序
// param           {*contextx.Contextx} ctx
// param           {[]model.ClassifyModel} cmList
// return          {*}
//
func ClassifyOrderUpdate(ctx *contextx.Contextx, cmList []model.ClassifyModel) (err error) {
	for i := 0; i < len(cmList); i++ {
		err = ctx.BaseTx.ClassifyModel.Update(cmList[i])
		if err != nil {
			return
		}
	}
	return
}

func ClassifyDel(ctx *contextx.Contextx, classifyId string) (err error) {
	result, err := ctx.BaseTx.TaskModel.FindByClassifyId(classifyId)
	if err != nil {
		return
	}
	if len(result) > 0 {
		err = i18n.NewCodeError(module.ClassifyDelExistTask)
		return
	}
	err = ctx.BaseTx.ClassifyModel.Delete(ctx.TID, classifyId)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("oc", classifyId), zap.String("tid", ctx.TID))
		err = i18n.NewCodeError(module.ClassifyUpdateErr)
	}
	return
}
