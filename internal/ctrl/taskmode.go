package ctrl

import (
	"net/http"

	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/model"
	"github.com/clz.skywalker/event.shop/kernal/internal/service"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreateTaskMode(c *gin.Context) {
	ret := httpx.NewResult()
	defer c.JSON(http.StatusOK, ret)
	tm := model.TaskModeModel{}
	err := c.ShouldBind(&tm)
	if err != nil {
		errx := i18n.NewCodeError(module.RequestParamBindCode, err.Error())
		utils.ZapLog.Error(errx.Msg)
		ret.SetCodeErr(errx)
		return
	}
	locale := c.GetHeader("Accept-Language")
	err = container.GlobalServerContext.Validator.ValidateParam(locale, &tm)
	if err != nil {
		errx := i18n.NewCodeError(module.TranslatorNotFoundCode, err.Error())
		utils.ZapLog.Error(errx.Msg)
		ret.SetCodeErr(errx)
		return
	}
	_, err = service.InsertTaskMode(container.GlobalServerContext.TaskModeModel, tm)
	if err != nil {
		ret.SetCodeErr(err)
		return
	}
	errx := i18n.NewCodeError(module.SuccessCode)
	ret.SetCodeErr(errx)
}
