package ctrl

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/gin-gonic/gin"
)

func validateBind(c *gin.Context, m interface{}) (ctx *contextx.Contextx, err error) {
	ctx = &contextx.Contextx{}
	lang := c.GetHeader(constx.HeaderLang)
	if m != nil {
		err = c.ShouldBind(&m)
		if err != nil && m != nil {
			err = i18n.NewCodeError(lang, module.RequestParamBindErr, err.Error())
			loggerx.ZapLog.Error(err.Error())
			return
		}
		err = container.GlobalServerContext.Validator.ValidateParam(lang, m)
		if err != nil && m != nil {
			err = i18n.NewCodeError(lang, module.TranslatorNotFoundErr, err.Error())
			loggerx.ZapLog.Error(err.Error())
			return
		}
	}

	ctx.Language = lang
	ctx.Context = c
	ctx.UID = c.GetString(constx.TokenUID)
	ctx.TID = c.GetString(constx.TokenTID)
	ctx.BaseTx = *container.GlobalServerContext
	return
}
