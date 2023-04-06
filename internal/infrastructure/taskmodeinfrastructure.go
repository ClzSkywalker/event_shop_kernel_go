package infrastructure

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
)

func InsertTaskMode(ctx *contextx.Contextx, tm *model.TaskModeModel) (id uint, err error) {
	id, err = ctx.BaseTx.TaskModeModel.Insert(tm)
	return
}
