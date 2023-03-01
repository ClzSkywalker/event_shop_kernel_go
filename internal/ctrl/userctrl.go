package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
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
