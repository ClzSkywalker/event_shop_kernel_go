package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/gin-gonic/gin"
)

func ClassifyQueryTeam(c *gin.Context) {
	ret := httpx.NewResult()
	var err error
	defer func() {
		c.JSON(http.StatusOK, ret)
	}()
	req := entity.ClassifyInsertReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	cms, err := infrastructure.ClassifyFindByTeamId(ctx)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	resp := entity.ClassifyFind{Data: make([]entity.ClassifyItem, 0, len(cms))}
	for i := 0; i < len(cms); i++ {
		resp.Data = append(resp.Data, entity.ClassifyItem{
			OnlyCode:  cms[i].OnlyCode,
			CreatedBy: cms[i].CreatedBy,
			Title:     cms[i].Title,
			Color:     cms[i].Color,
			Sort:      cms[i].Sort,
		})
	}
	ret.Data = resp.Data
}

func ClassifyInsert(c *gin.Context) {
	ret := httpx.NewResult()
	var err error
	defer func() {
		c.JSON(http.StatusOK, ret)
	}()
	req := entity.ClassifyInsertReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	cid, err := service.ClassifyCreate(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = entity.CommonResponseId{OnlyCode: cid}
}

func ClassifyUpdate(c *gin.Context) {
	ret := httpx.NewResult()
	var err error
	defer func() {
		c.JSON(http.StatusOK, ret)
	}()
	req := entity.ClassifyItem{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	err = service.ClassifyUpdate(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}

func ClassifyOrderUpdate(c *gin.Context) {
	ret := httpx.NewResult()
	var err error
	defer func() {
		c.JSON(http.StatusOK, ret)
	}()
	req := entity.ClassifyOrderReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	err = service.ClassifyOrderUpdate(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}

func ClassifyDel(c *gin.Context) {
	ret := httpx.NewResult()
	var err error
	defer func() {
		c.JSON(http.StatusOK, ret)
	}()
	req := entity.ClassifyDelReq{}
	ctx, err := validateBind(c, &req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	err = service.ClassifyDel(ctx, req)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}
