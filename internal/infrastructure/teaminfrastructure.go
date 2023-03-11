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

func TeamQueryAllByCreatedBy(ctx contextx.Contextx, tm model.ITeamModel) (tmList []model.TeamModel, err error) {
	tmList, err = tm.Find(model.TeamModel{CreatedBy: ctx.UID})
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(ctx.Language, module.TeamFindErr)
		return
	}
	return
}

func TeamFindMyTeam(ctx *contextx.Contextx) (tmList []entity.TeamItem, err error) {
	tm := model.NewDefaultTeamModel(ctx.Tx)
	tmList, err = tm.FindMyTeam(ctx.UID)
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.TeamFindErr)
		return
	}
	return
}

func TeamQueryFirst(ctx *contextx.Contextx, p model.TeamModel) (result model.TeamModel, err error) {
	tm := model.NewDefaultTeamModel(ctx.Tx)
	result, err = tm.First(p)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(ctx.Language, module.TeamFindErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(ctx.Language, module.TeamFindErr)
		return
	}
	return
}

func TeamCreate(ctx *contextx.Contextx, tm *model.TeamModel, sort int) (tid string, err error) {
	tx := model.NewDefaultTeamModel(ctx.Tx)
	utm := model.NewDefaultUserToTeamModel(ctx.Tx)
	tid = utils.NewUlid()
	tm.TeamId = tid
	tm.CreatedBy = ctx.TID
	_, err = tx.Create(tm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = i18n.NewCodeError(ctx.Language, module.TeamCreateErr)
		return
	}
	_, err = utm.Insert(&model.UserToTeamModel{Uid: ctx.UID, Tid: tid, Sort: sort})
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = i18n.NewCodeError(ctx.Language, module.TeamCreateErr)
		return
	}
	return
}

func TeamUpdate(ctx *contextx.Contextx, m model.TeamModel) (err error) {
	tm := model.NewDefaultTeamModel(ctx.Tx)
	_, err = tm.First(model.TeamModel{TeamId: m.TeamId})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.TeamNotFound)
		return
	}
	err = tm.Update(m)
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.TeamUpdateErr)
		return
	}
	return
}

func TeamDelete(ctx *contextx.Contextx, teamId string) (err error) {
	tm := model.NewDefaultTeamModel(ctx.Tx)
	_, err = tm.First(model.TeamModel{TeamId: teamId, CreatedBy: teamId})
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(ctx.Language, module.TeamFindErr)
		return
	}
	err = tm.Delete(teamId)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(ctx.Language, module.TeamDeleteErr)
		return
	}
	return
}
