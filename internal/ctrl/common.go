package ctrl

import (
	"github.com/clz.skywalker/event.shop/kernal/internal/container"
	"github.com/clz.skywalker/event.shop/kernal/internal/contextx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
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
			err = errorx.NewCodeError(module.RequestParamBindErr, err.Error())
			loggerx.ZapLog.Error(err.Error())
			return
		}
		err = container.GlobalServerContext.Validator.ValidateParam(lang, m)
		if err != nil && m != nil {
			err = errorx.NewCodeError(module.TranslatorNotFoundErr, err.Error())
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

func getResult(c *gin.Context) *httpx.Result {
	lang := c.GetHeader(constx.HeaderLang)
	switch lang {
	case constx.LangChinese, constx.LangEnglish:
	default:
		lang = constx.LangDefault
	}
	ret := httpx.NewResult()
	if lang == "" {
		lang = constx.LangDefault
	}
	ret.Lang = lang
	return ret
}
