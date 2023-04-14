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

func UserRegisterByEmail(ctx *contextx.Contextx, req entity.RegisterByEmailReq) (token string, err error) {
	um := model.UserModel{}
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		um, err = infrastructure.RegisterByEmail(ctx, req)
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
		err = errorx.NewCodeError(module.UserRegisterErr)
		return
	}
	return
}

func UserRegisterByPhone(ctx *contextx.Contextx, req entity.RegisterByPhoneReq) (token string, err error) {
	um := model.UserModel{}
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		um, err = infrastructure.RegisterByPhone(ctx, req)
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
		err = errorx.NewCodeError(module.UserRegisterErr)
		return
	}
	return
}

func UserRegisterByUid(ctx *contextx.Contextx) (token string, err error) {
	um := model.UserModel{}
	err = ctx.BaseTx.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		um, err = infrastructure.RegisterByUid(ctx)
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
		err = errorx.NewCodeError(module.UserRegisterErr)
		return
	}
	return
}

func UserResetPwd(ctx *contextx.Contextx, req entity.UserResetPwdReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err = infrastructure.UserResetPwd(ctx, req.OldPwd, req.OldPwd)
		return err
	})
	return
}

func UserUpdate(ctx *contextx.Contextx, req entity.UserUpdateReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err = infrastructure.UserUpdate(ctx, model.UserModel{
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

func UserBindEmail(ctx *contextx.Contextx, req entity.BindEmailReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err = infrastructure.BindEmailByUid(ctx, req)
		return err
	})
	return
}

func UserBindPhone(ctx *contextx.Contextx, req entity.BindPhoneReq) (err error) {
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		ctx.BaseTx = *container.NewBaseServiceContext(&ctx.BaseTx, tx)
		err = infrastructure.BindPhoneByUid(ctx, req)
		return err
	})
	return
}
