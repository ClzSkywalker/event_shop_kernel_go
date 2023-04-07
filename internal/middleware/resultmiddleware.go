package middleware

import (
	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/gin-gonic/gin"
)

func ResultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
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
		c.Set(constx.CtxRet, ret)
		c.Next()
	}
}
