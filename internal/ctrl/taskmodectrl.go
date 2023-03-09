package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/gin-gonic/gin"
)

func CreateTaskMode(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	tm := model.TaskModeModel{}
	ctx, err := validateBind(c, &tm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	var id uint
	id, err = service.InsertTaskMode(container.GlobalServerContext.TaskModeModel, &tm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	errx := i18n.NewCodeError(ctx.Language, module.SuccessCode)
	ret.Data = id
	ret.SetCodeErr(errx)
}
