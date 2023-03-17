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

func UserRegisterByEmail(ctx *contextx.Contextx, req entity.RegisterByEmailReq) (token string, err error) {
	um := model.UserModel{}
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		base := &container.BaseServiceContext{}
		base = container.NewBaskServiceContext(base, tx)
		um, err = infrastructure.RegisterByEmail(ctx, base.UserModel, base.TeamModel, req)
		return err
	})
	if err != nil {
		return
	}
	token, err = infrastructure.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterErr)
		return
	}
	return
}

func UserRegisterByPhone(ctx *contextx.Contextx, req entity.RegisterByPhoneReq) (token string, err error) {
	um := model.UserModel{}
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		base := &container.BaseServiceContext{}
		base = container.NewBaskServiceContext(base, tx)
		um, err = infrastructure.RegisterByPhone(ctx, base.UserModel, base.TeamModel, req)
		return err
	})
	if err != nil {
		return
	}

	token, err = infrastructure.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterErr)
		return
	}
	return
}

func UserRegisterByUid(ctx *contextx.Contextx) (token string, err error) {
	um := model.UserModel{}
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		base := &container.BaseServiceContext{}
		base = container.NewBaskServiceContext(base, tx)
		um, err = infrastructure.RegisterByUid(ctx, base.UserModel, base.TeamModel)
		return err
	})
	token, err = infrastructure.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterErr)
		return
	}
	return
}

func UserUpdate(ctx *contextx.Contextx, req entity.UserUpdateReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		base := &container.BaseServiceContext{}
		base = container.NewBaskServiceContext(base, tx)
		err = infrastructure.UserUpdate(ctx, base.UserModel, model.UserModel{
			CreatedBy:  req.CreatedBy,
			TeamIdPort: req.TeamIdPort,
			NickName:   req.NickName,
			Phone:      req.Phone,
			Version:    req.Version,
		})
		return err
	})
	return
}
