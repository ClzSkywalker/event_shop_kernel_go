package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/infrastructure"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/gin-gonic/gin"
)

func CreateTaskMode(c *gin.Context) {
	ret := getResult(c)
	defer c.JSON(http.StatusOK, ret)
	tm := model.TaskModeModel{}
	ctx, err := validateBind(c, &tm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	var id uint
	id, err = infrastructure.InsertTaskMode(ctx, &tm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	errx := i18n.NewCodeError(module.SuccessCode)
	ret.Data = id
	ret.SetCodeErr(errx)
}
