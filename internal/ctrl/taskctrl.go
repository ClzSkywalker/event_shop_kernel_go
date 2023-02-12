package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/gin-gonic/gin"
)

func InsertTask(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	tm := model.TaskModel{}
	err := validateBind(c, &tm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	id, err := service.InsertTask(container.GlobalServerContext.TaskModel, &tm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	ret.Data = id
}