package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/gin-gonic/gin"
)

func TeamCreate(c *gin.Context) {
	ret := getResult(c)
	var err error
	defer func() {
		c.JSON(http.StatusOK, ret)
	}()
	cm := entity.TeamCreateReq{}
	ctx, err := validateBind(c, &cm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	tid, err := service.TeamCreate(ctx, cm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.TeamCreateResp{TeamId: tid}
}

func TeamUpdate(c *gin.Context) {
	ret := getResult(c)
	var err error
	defer func() {
		c.JSON(http.StatusOK, ret)
	}()
	cm := entity.TeamUpdateReq{}
	ctx, err := validateBind(c, &cm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	err = service.TeamUpdate(ctx, cm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}

func TeamDel(c *gin.Context) {
	ret := getResult(c)
	var err error
	defer func() {
		c.JSON(http.StatusOK, ret)
	}()
	cm := entity.TeamDelReq{}
	ctx, err := validateBind(c, &cm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	err = service.TeamDel(ctx, cm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}

func TeamFindMyTeam(c *gin.Context) {
	ret := getResult(c)
	var err error
	defer func() {
		c.JSON(http.StatusOK, ret)
	}()
	ctx, err := validateBind(c, nil)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	tmList, err := infrastructure.TeamFindMyTeam(ctx)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = tmList
}
