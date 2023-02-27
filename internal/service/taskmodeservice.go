package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
)

func InsertTaskMode(tx model.ITaskModeModel, tm *model.TaskModeModel) (id uint, err error) {
	id, err = tx.Insert(tm)
	return
}
