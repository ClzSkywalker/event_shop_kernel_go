package ctrl

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/gin-gonic/gin"
)

func validateBind(c *gin.Context, m interface{}) (err error) {
	err = c.ShouldBind(&m)
	if err != nil {
		errx := i18n.NewCodeError(module.RequestParamBindErr, err.Error())
		loggerx.ZapLog.Error(errx.Msg)
		return errx
	}
	locale := c.GetHeader("Accept-Language")
	err = container.GlobalServerContext.Validator.ValidateParam(locale, m)
	if err != nil {
		errx := i18n.NewCodeError(module.TranslatorNotFoundErr, err.Error())
		loggerx.ZapLog.Error(errx.Msg)
		return errx
	}
	return
}
