package service

import "github.com/clz.skywalker/event.shop/kernal/internal/model"

func InsertTaskMode(m model.ITaskModeModel, tm model.TaskModeModel) (id int64, err error) {
	return m.Insert(tm)
}
