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

func TeamQueryAllByCreatedBy(ctx contextx.Contextx) (tmList []model.TeamModel, err error) {
	tmList, err = ctx.BaseTx.TeamModel.Find(model.TeamModel{CreatedBy: ctx.UID})
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(module.TeamFindErr)
		return
	}
	return
}

func TeamFindMyTeam(ctx *contextx.Contextx) (tmList []model.TeamModel, err error) {
	tmList, err = ctx.BaseTx.TeamModel.FindMyTeam(ctx.UID)
	if err != nil {
		err = i18n.NewCodeError(module.TeamFindErr)
		return
	}
	return
}

func TeamQueryFirst(ctx *contextx.Contextx, p model.TeamModel) (result model.TeamModel, err error) {
	result, err = ctx.BaseTx.TeamModel.First(p)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(module.TeamFindErr)
		return
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = i18n.NewCodeError(module.TeamFindErr)
		return
	}
	return
}

func TeamCreate(ctx *contextx.Contextx, tm *model.TeamModel, sort int) (tid string, err error) {
	team, err := ctx.BaseTx.TeamModel.FindMyTeamName(ctx.UID, tm.Name)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		return
	}
	if team.Id != 0 {
		err = i18n.NewCodeError(module.TeamNameRepeatErr)
		return
	}

	tid = utils.NewUlid()
	tm.TeamId = tid
	tm.CreatedBy = ctx.TID
	_, err = ctx.BaseTx.TeamModel.Create(tm)
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = i18n.NewCodeError(module.TeamCreateErr)
		return
	}
	_, err = ctx.BaseTx.UserToTeamModel.Insert(&model.UserToTeamModel{Uid: ctx.UID, Tid: tid, Sort: sort})
	if err != nil {
		loggerx.ZapLog.Error(err.Error(), zap.Any("model", tm))
		err = i18n.NewCodeError(module.TeamCreateErr)
		return
	}
	return
}

func TeamUpdate(ctx *contextx.Contextx, m model.TeamModel) (err error) {
	_, err = ctx.BaseTx.TeamModel.First(model.TeamModel{TeamId: m.TeamId})
	if err != nil {
		err = i18n.NewCodeError(module.TeamNotFound)
		return
	}
	err = ctx.BaseTx.TeamModel.Update(m)
	if err != nil {
		err = i18n.NewCodeError(module.TeamUpdateErr)
		return
	}
	return
}

func TeamDelete(ctx *contextx.Contextx, teamId string) (err error) {
	_, err = ctx.BaseTx.TeamModel.First(model.TeamModel{TeamId: teamId, CreatedBy: teamId})
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(module.TeamFindErr)
		return
	}
	err = ctx.BaseTx.TeamModel.Delete(teamId)
	if err != nil {
		loggerx.ZapLog.Error(err.Error())
		err = i18n.NewCodeError(module.TeamDeleteErr)
		return
	}
	return
}
