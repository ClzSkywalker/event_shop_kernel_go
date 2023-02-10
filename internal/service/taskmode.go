package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"go.uber.org/zap"
)

func InsertTaskMode(m model.ITaskModeModel, tm model.TaskModeModel) (id int64, err error) {
	id, err = m.Insert(tm)
	if err != nil {
		utils.ZapLog.Error(err.Error(), zap.Any("model", tm))
	}
	err = i18n.NewCodeError(module.TaskModeErr, err)
	return
}
