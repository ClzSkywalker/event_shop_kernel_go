package service

import (
	"errors"

	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
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

func TeamFindMyTeam(ctx *contextx.Contextx, tm model.ITeamModel) (tmList []model.TeamModel, err error) {
	tmList, err = tm.FindMyTeam(ctx.UID)
	return
}

func TeamQueryFirst(ctx contextx.Contextx, tm model.ITeamModel, p model.TeamModel) (result model.TeamModel, err error) {
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

func TeamUpdate(ctx contextx.Contextx, tm model.ITeamModel, m model.TeamModel) (err error) {
	_, err = tm.First(model.TeamModel{TeamId: m.TeamId})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.TeamFindErr)
		return
	}
	err = tm.Update(m)
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.TeamUpdateErr)
		return
	}
	return
}

func TeamDelete(ctx contextx.Contextx, tm model.ITeamModel, teamId string) (err error) {
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
