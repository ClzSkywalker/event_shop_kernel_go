package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/gin-gonic/gin"
)

func RegisterByEmail(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	req := entity.RegisterByEmailReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.UserRegisterByEmail(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func RegisterByPhone(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	req := entity.RegisterByPhoneReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.UserRegisterByPhone(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func RegisterByUid(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	ctx, err := validateBind(c, nil)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}

	token, err := service.UserRegisterByUid(ctx)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}

	ret.Data = entity.LoginRep{Token: token}
}

func LoginByEmail(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	req := entity.LoginByEmailReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	um, err := infrastructure.LoginByEmail(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := infrastructure.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = errorx.NewCodeError(module.UserRegisterErr)
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
	um, err := infrastructure.LoginByPhone(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := infrastructure.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = errorx.NewCodeError(module.UserRegisterErr)
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
	um, err := infrastructure.LoginByUid(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := infrastructure.GenerateToken(entity.TokenInfo{
		UID: um.CreatedBy,
		TID: um.TeamIdPort,
	})
	if err != nil {
		err = errorx.NewCodeError(module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func UserGetInfo(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	ctx, err := validateBind(c, nil)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	um, err := infrastructure.GetUserInfo(ctx, ctx.UID)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	resp := entity.UserResp{
		UserItem: entity.UserItem{
			CreatedBy:    um.CreatedBy,
			TeamIdPort:   um.TeamIdPort,
			NickName:     um.NickName,
			MemberType:   um.MemberType,
			RegisterType: um.RegisterType,
			Picture:      um.Picture,
			Email:        um.Email,
			Phone:        um.Phone,
			Version:      um.Version,
		},
	}
	ret.Data = resp
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
	err = infrastructure.BindEmailByUid(ctx, req)
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
	err = infrastructure.BindPhoneByUid(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}

func UserUpdate(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.UserUpdateReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	err = service.UserUpdate(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}
