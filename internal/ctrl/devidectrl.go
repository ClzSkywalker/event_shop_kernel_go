package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/gin-gonic/gin"
)

func DevideInsert(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	req := entity.DevideInsertReqEntity{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	oc, err := service.DevideInsert(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.CommonResponseId{OnlyCode: oc}
}

func DevideUpdate(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	req := entity.DevideUpdateReqEntity{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	err = service.DevideUpdate(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}

func DevideDelete(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	req := entity.CommonRequestId{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	err = service.DevideDelete(ctx, req.OnlyCode)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}
