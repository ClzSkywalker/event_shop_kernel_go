package middleware

import (
	"net/http"
	"time"

	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/gin-gonic/gin"
)

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-03
 * @Description    : jwt中间件，用于处理登录用户接口
 * @return          {*}
 */
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		ret := httpx.NewResult()
		if token == "" {
			err := i18n.NewCodeError(module.TokenInvalid)
			ret.SetCodeErr(err)
			c.JSON(http.StatusOK, ret)
			c.Abort()
		}
		t, err := utils.ParseToken(token, constx.TokenSecret)
		if err != nil {
			ret.SetCodeErr(err)
			c.JSON(http.StatusOK, ret)
			c.Abort()
		}
		epTime, err := t.Claims.GetExpirationTime()
		if err != nil {
			ret.SetCodeErr(err)
			c.JSON(http.StatusOK, ret)
			c.Abort()
		}
		if epTime.Time.Before(time.Now()) {
			err := i18n.NewCodeError(module.TokenExpired)
			ret.SetCodeErr(err)
			c.JSON(http.StatusOK, ret)
			c.Abort()
		}
		c.Next()
	}
}
