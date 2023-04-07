package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"gorm.io/gorm"
)

func TaskInsert(ctx *contextx.Contextx, req entity.TaskEntity) (oc string, err error) {
	m := model.TaskModel{}
	err = StructToStruct(ctx, req, &m)
	if err != nil {
		return
	}

	err = ctx.BaseTx.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		oc, err = infrastructure.TaskInsert(ctx, &m)
		return err
	})
	return
}
