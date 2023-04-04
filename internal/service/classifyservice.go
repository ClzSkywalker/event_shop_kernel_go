package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"gorm.io/gorm"
)

func ClassifyCreate(ctx *contextx.Contextx, req entity.ClassifyInsertReq) (cid string, err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		cid, err = infrastructure.ClassifyInsert(ctx, container.GlobalServerContext.ClassifyModel, &model.ClassifyModel{
			CreatedBy: ctx.UID,
			TeamId:    ctx.TID,
			Title:     req.Title,
			Color:     req.Color,
			Sort:      req.Sort,
		})
		return err
	})
	return
}

func ClassifyUpdate(ctx *contextx.Contextx, req entity.ClassifyItem) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		err = infrastructure.ClassifyUpdate(ctx, container.GlobalServerContext.ClassifyModel, model.ClassifyModel{
			CreatedBy: ctx.UID,
			TeamId:    ctx.TID,
			OnlyCode:  req.OnlyCode,
			Title:     req.Title,
			Color:     req.Color,
			Sort:      req.Sort,
		})
		return err
	})
	return
}

func ClassifyDel(ctx *contextx.Contextx, req entity.ClassifyDelReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		err = infrastructure.ClassifyDel(ctx, container.GlobalServerContext.ClassifyModel, req.OnlyCode)
		return err
	})
	return
}
