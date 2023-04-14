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

func TaskInsert(ctx *contextx.Contextx, req entity.TaskInsertReq) (oc string, err error) {
	m := model.TaskModel{
		Title:      req.Title,
		DevideId:   req.DevideId,
		TaskModeId: req.TaskModeId,
		StartAt:    req.StartAt,
		EndAt:      req.EndAt,
		ParentId:   req.ParentId,
	}

	err = ctx.BaseTx.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		coc, err := infrastructure.TaskContentInsert(ctx, req.Content)
		if err != nil {
			return err
		}
		m.ContentId = coc
		oc, err = infrastructure.TaskInsert(ctx, &m)
		return err
	})
	return
}

func TaskUpdate(ctx *contextx.Contextx, req entity.TaskUpdateReq) (err error) {
	m := model.TaskModel{
		OnlyCode:    req.OnlyCode,
		Title:       req.Title,
		DevideId:    req.DevideId,
		TaskModeId:  req.TaskModeId,
		CompletedAt: req.CompletedAt,
		GiveUpAt:    req.CompletedAt,
		StartAt:     req.StartAt,
		EndAt:       req.EndAt,
		ParentId:    req.ParentId,
	}
	err = StructToStruct(ctx, req, &m)
	if err != nil {
		return
	}

	err = ctx.BaseTx.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		tm, err := infrastructure.TaskFirst(ctx, model.TaskModel{OnlyCode: m.OnlyCode})
		if err != nil {
			return err
		}
		if tm.CreatedBy != ctx.UID {
			err = errorx.NewCodeError(module.OperateNoPermission)
			return err
		}
		m.ContentId = tm.ContentId
		err = infrastructure.TaskUpdate(ctx, m)
		return err
	})
	return
}

func TaskDelete(ctx *contextx.Contextx, oc string) (err error) {
	err = ctx.BaseTx.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err = infrastructure.TaskDelete(ctx, oc)
		return err
	})
	return
}
