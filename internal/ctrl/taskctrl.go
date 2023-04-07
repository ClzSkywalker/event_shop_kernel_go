package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/entity"
	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/gin-gonic/gin"
)

func FindClassifyById(c *gin.Context) {
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
	var data []entity.TaskEntity
	err = utils.StructToStruct(taskList, &data)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = data
}

func TaskInsert(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	tm := entity.TaskEntity{}
	ctx, err := validateBind(c, &tm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	id, err := service.TaskInsert(ctx, tm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = id
}
