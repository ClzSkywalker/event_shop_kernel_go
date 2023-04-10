package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/gin-gonic/gin"
)

func TaskFilter(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	param := entity.TaskFilterParam{}
	ctx, err := validateBind(c, &param)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	result, err := infrastructure.TaskFilter(ctx, param)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = result
}

func TaskFindByClassifyId(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	param := entity.TaskFindByClassifyIdEntity{}
	ctx, err := validateBind(c, &param)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	taskList, err := infrastructure.TaskFindByClassifyId(ctx, param.ClassifyId)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = taskList
}

func TaskInsert(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	param := entity.TaskInsertReq{}
	ctx, err := validateBind(c, &param)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	id, err := service.TaskInsert(ctx, param)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = id
}

func TaskUpdate(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	param := entity.TaskUpdateReq{}
	ctx, err := validateBind(c, &param)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	err = service.TaskUpdate(ctx, param)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}

func TaskDelete(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	param := entity.TaskDeleteReq{}
	ctx, err := validateBind(c, &param)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	err = service.TaskDelete(ctx, param.OnlyCode)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
}
