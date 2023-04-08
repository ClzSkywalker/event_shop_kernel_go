package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"gorm.io/gorm"
)

func ClassifyCreate(ctx *contextx.Contextx, req entity.ClassifyInsertReq) (cid string, err error) {
	m := model.ClassifyModel{}
	err = StructToStruct(ctx, req, &m)
	if err != nil {
		return
	}
	m.CreatedBy = ctx.UID
	m.TeamId = ctx.TID

	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		cid, err = infrastructure.ClassifyInsert(ctx, &m)
		return err
	})
	return
}

func ClassifyUpdate(ctx *contextx.Contextx, req entity.ClassifyUpdateReq) (err error) {
	m := model.ClassifyModel{}
	err = StructToStruct(ctx, req, &m)
	if err != nil {
		return
	}

	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err = infrastructure.ClassifyUpdate(ctx, m)
		return err
	})
	return
}

func ClassifyOrderUpdate(ctx *contextx.Contextx, req entity.ClassifyOrderReq) (err error) {
	cmList := make([]model.ClassifyModel, 0, len(req.Data))
	for _, item := range req.Data {
		if item.OnlyCode == "" {
			err = i18n.NewCodeError(module.ReqMissErr)
			return
		}
		cmList = append(cmList, model.ClassifyModel{
			OnlyCode: item.OnlyCode,
			Sort:     item.Sort,
		})
	}
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err = infrastructure.ClassifyOrderUpdate(ctx, cmList)
		return err
	})
	return
}

func ClassifyDel(ctx *contextx.Contextx, req entity.ClassifyDelReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err = infrastructure.ClassifyDel(ctx, req.OnlyCode)
		return err
	})
	return
}
