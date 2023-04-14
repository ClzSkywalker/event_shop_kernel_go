package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"gorm.io/gorm"
)

func DevideInsert(ctx *contextx.Contextx, req entity.DevideInsertReqEntity) (oc string, err error) {
	m := model.DevideModel{
		Title:      req.Title,
		Sort:       req.Sort,
		ClassifyId: req.ClassifyId,
	}
	err = ctx.BaseTx.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		oc, err = infrastructure.DevideInsert(ctx, m)
		return err
	})
	return
}

func DevideUpdate(ctx *contextx.Contextx, req entity.DevideUpdateReqEntity) (err error) {
	m := model.DevideModel{
		OnlyCode:   req.OnlyCode,
		Title:      req.Title,
		Sort:       req.Sort,
		ClassifyId: req.ClassifyId,
	}
	err = ctx.BaseTx.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err = infrastructure.DevideUpdate(ctx, m)
		return err
	})
	return
}

func DevideDelete(ctx *contextx.Contextx, oc string) (err error) {
	err = ctx.BaseTx.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		_, err := infrastructure.TaskFirst(ctx, model.TaskModel{DevideId: oc})
		if err != nil && !errorx.Is(err, module.TaskNotfoundErr) {
			err = errorx.NewCodeError(module.DevideDelExistTaskErr)
			return err
		}
		err = infrastructure.DevideDelete(ctx, oc)
		return err
	})
	return
}
