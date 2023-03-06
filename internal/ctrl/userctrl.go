package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/gin-gonic/gin"
)

func RegisterByEmail(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.RegisterByEmailReq{}
	err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	uid, err := service.RegisterByEmail(container.GlobalServerContext.UserModel, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.GenerateToken(uid)
	if err != nil {
		err = i18n.NewCodeError(module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func RegisterByPhone(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.RegisterByPhoneReq{}
	err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	uid, err := service.RegisterByPhone(container.GlobalServerContext.UserModel, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.GenerateToken(uid)
	if err != nil {
		err = i18n.NewCodeError(module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func RegisterByUid(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.RegisterByPhoneReq{}
	err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	uid, err := service.RegisterByUid(container.GlobalServerContext.UserModel)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.GenerateToken(uid)
	if err != nil {
		err = i18n.NewCodeError(module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.RegisterByUidRep{LoginRep: entity.LoginRep{Token: token}, Uid: uid}
}

func LoginByEmail(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.LoginByEmailReq{}
	err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	uid, err := service.LoginByEmail(container.GlobalServerContext.UserModel, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.GenerateToken(uid)
	if err != nil {
		err = i18n.NewCodeError(module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func LoginByPhone(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.LoginByPhoneReq{}
	err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	uid, err := service.LoginByPhone(container.GlobalServerContext.UserModel, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.GenerateToken(uid)
	if err != nil {
		err = i18n.NewCodeError(module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func LoginByUid(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.LoginByUidReq{}
	err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	uid, err := service.LoginByUid(container.GlobalServerContext.UserModel, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	token, err := service.GenerateToken(uid)
	if err != nil {
		err = i18n.NewCodeError(module.UserRegisterErr)
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.LoginRep{Token: token}
}

func BindEmailByUid(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.BindEmailReq{}
	err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	uid, _ := c.Get(constx.TokenUid)
	err = service.BindEmailByUid(container.GlobalServerContext.UserModel, uid.(string), req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}

func BindPhoneByUid(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	req := entity.BindPhoneReq{}
	err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	uid, _ := c.Get(constx.TokenUid)
	err = service.BindPhoneByUid(container.GlobalServerContext.UserModel, uid.(string), req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}
