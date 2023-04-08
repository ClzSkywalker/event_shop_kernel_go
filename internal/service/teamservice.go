package service

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"gorm.io/gorm"
)

func TeamCreate(ctx *contextx.Contextx, req entity.TeamCreateReq) (tid string, err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		tid, err = infrastructure.TeamCreate(ctx, &model.TeamModel{
			Name:        req.Name,
			Description: req.Description,
		}, req.Sort)
		cm := model.NewDefaultClassifyModel(tx)
		cid := utils.NewUlid()
		err = cm.InitData(ctx.Language, ctx.UID, tid, cid)
		return err
	})
	return
}

func TeamUpdate(ctx *contextx.Contextx, req entity.TeamUpdateReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err := infrastructure.TeamUpdate(ctx, model.TeamModel{Name: req.Name, Description: req.Description,
			OnlyCode: req.TeamId})
		return err
	})
	return
}

func TeamDel(ctx *contextx.Contextx, req entity.TeamDelReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err := infrastructure.TeamDelete(ctx, req.TeamId)
		return err
	})
	return
}
