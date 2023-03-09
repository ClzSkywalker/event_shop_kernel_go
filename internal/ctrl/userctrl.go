package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterByEmail(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.RegisterByEmailReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	um := model.UserModel{}
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		base := &container.BaseServiceContext{}
		base = container.NewBaskServiceContext(base, tx)
		um, err = service.RegisterByEmail(ctx, base.UserModel, base.TeamModel, req)
		return err
	})
	if err != nil {
		ret.SetCodeErr(err)
		return
	}

	token, err := service.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func RegisterByPhone(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.RegisterByPhoneReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	um := model.UserModel{}
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		base := &container.BaseServiceContext{}
		base = container.NewBaskServiceContext(base, tx)
		um, err = service.RegisterByPhone(ctx, base.UserModel, base.TeamModel, req)
		return err
	})
	if err != nil {
		ret.SetCodeErr(err)
		return
	}

	token, err := service.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func RegisterByUid(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	ctx, err := validateBind(c, nil)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}

	um := model.UserModel{}
	err = container.GlobalServerContext.Db.Transaction(func(tx *gorm.DB) error {
		base := &container.BaseServiceContext{}
		base = container.NewBaskServiceContext(base, tx)
		um, err = service.RegisterByUid(ctx, base.UserModel, base.TeamModel)
		return err
	})
	if err != nil {
		ret.SetCodeErr(err)
		return
	}

	token, err := service.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.RegisterByUidRep{LoginRep: entity.LoginRep{Token: token}, Uid: um.CreatedBy}
}

func LoginByEmail(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.LoginByEmailReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	um, err := service.LoginByEmail(ctx, container.GlobalServerContext.UserModel, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func LoginByPhone(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.LoginByPhoneReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	um, err := service.LoginByPhone(ctx, container.GlobalServerContext.UserModel, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func LoginByUid(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.LoginByUidReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	um, err := service.LoginByUid(ctx, container.GlobalServerContext.UserModel, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = i18n.NewCodeError(ctx.Language, module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func BindEmailByUid(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.BindEmailReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	uid, _ := c.Get(constx.TokenUID)
	err = service.BindEmailByUid(ctx, container.GlobalServerContext.UserModel, uid.(string), req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}

func BindPhoneByUid(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.BindPhoneReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	uid, _ := c.Get(constx.TokenUID)
	err = service.BindPhoneByUid(ctx, container.GlobalServerContext.UserModel, uid.(string), req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}
