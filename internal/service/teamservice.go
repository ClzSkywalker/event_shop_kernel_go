package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"gorm.io/gorm"
)

func TeamCreate(ctx *contextx.Contextx, req entity.TeamCreateReq) (tid string, err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.Tx = tx
		tid, err = infrastructure.TeamCreate(ctx, &model.TeamModel{
			Name:        req.Name,
			Description: req.Description,
		}, req.Sort)
		return err
	})
	return
}

func TeamUpdate(ctx *contextx.Contextx, req entity.TeamUpdateReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.Tx = tx
		err := infrastructure.TeamUpdate(ctx, model.TeamModel{Name: req.Name, Description: req.Description,
			TeamId: req.TeamId})
		return err
	})
	return
}

func TeamDel(ctx *contextx.Contextx, req entity.TeamDelReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.Tx = tx
		err := infrastructure.TeamDelete(ctx, req.TeamId)
		return err
	})
	return
}
