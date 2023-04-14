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

func ClassifyFindByTeamId(ctx *contextx.Contextx) (cms []model.ClassifyModel, err error) {
	cms, err = ctx.BaseTx.ClassifyModel.FindByTeamId(ctx.TID)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("ctx", ctx))
		err = errorx.NewCodeError(module.ClassifyQueryErr)
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
		err = errorx.NewCodeError(module.ClassifyNotfoundErr)
		return
	}
	err = errorx.NewCodeError(module.ClassifyQueryErr)
	loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
	return
}

func ClassifyInsert(ctx *contextx.Contextx, cm *model.ClassifyModel) (cid string, err error) {
	_, err = TeamFirst(ctx, model.TeamModel{OnlyCode: ctx.TID})
	if err != nil {
		return
	}

	if cm.ParentId != "" {
		var cp model.ClassifyModel
		cp, err = ClassifyFirst(ctx, model.ClassifyModel{OnlyCode: cm.ParentId})
		if err != nil {
			err = errorx.NewCodeError(module.ClassifyParentNoExistErr)
			return
		}
		if cp.ParentId != "" {
			err = errorx.NewCodeError(module.ClassifyDeepErr)
			return
		}
	}

	m, err := ClassifyFirst(ctx, model.ClassifyModel{TeamId: cm.TeamId, Title: cm.Title})

	if err != nil && !errorx.Is(err, module.ClassifyNotfoundErr) {
		return
	} else if err == nil && m.OnlyCode != cm.OnlyCode {
		err = errorx.NewCodeError(module.ClassifyTitleRepeatErr)
		return
	}

	oc := utils.NewUlid()
	cm.OnlyCode = oc
	cm.CreatedBy = ctx.UID
	cm.TeamId = ctx.TID
	_, err = ctx.BaseTx.ClassifyModel.Insert(cm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
		err = errorx.NewCodeError(module.ClassifyInsertErr)
	}
	cid = oc
	return
}

func ClassifyUpdate(ctx *contextx.Contextx, cm model.ClassifyModel) (err error) {
	_, err = TeamFirst(ctx, model.TeamModel{OnlyCode: ctx.TID})
	if err != nil {
		return
	}

	m, err := ClassifyFirst(ctx, model.ClassifyModel{TeamId: cm.TeamId, Title: cm.Title})

	if err != nil && !errorx.Is(err, module.ClassifyNotfoundErr) {
		return
	} else if err == nil && m.OnlyCode != cm.OnlyCode {
		err = errorx.NewCodeError(module.ClassifyTitleRepeatErr)
		return
	}

	err = ctx.BaseTx.ClassifyModel.Update(cm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", cm))
		err = errorx.NewCodeError(module.ClassifyUpdateErr)
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

func ClassifyDel(ctx *contextx.Contextx, oc string) (err error) {
	_, err = DevideFirst(ctx, model.DevideModel{ClassifyId: oc})
	if err == nil {
		err = errorx.NewCodeError(module.ClassifyDelExistDevideErr)
		return
	}

	err = ctx.BaseTx.ClassifyModel.Delete(ctx.TID, oc)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.String("oc", oc), zap.String("tid", ctx.TID))
		err = errorx.NewCodeError(module.ClassifyUpdateErr)
	}
	return
}
