package middleware

import (
	"net/http"
	"time"

	"github.com/clz.skywalker/event.shop/kernal/pkg/constx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/httpx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/errorx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/i18n/module"
	"github.com/clz.skywalker/event.shop/kernal/pkg/loggerx"
	"github.com/clz.skywalker/event.shop/kernal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

/**
 * @Author         : ClzSkywalker
 * @Date           : 2023-03-03
 * @Description    : jwt中间件，用于处理登录用户接口
 * @return          {*}
 */
func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(constx.HeaderAuthorization)
		value, _ := c.Get(constx.CtxRet)
		ret := value.(*httpx.Result)

		if token == "" {
			err := errorx.NewCodeError(module.TokenInvalid)
			ret.SetCodeErr(err)
			c.JSON(http.StatusOK, ret)
			c.Abort()
			return
		}
		t, err := utils.ParseToken(token, constx.TokenSecret)
		if err != nil {
			ret.SetCodeErr(err)
			c.JSON(http.StatusOK, ret)
			c.Abort()
			return
		}
		epTime, err := t.Claims.GetExpirationTime()
		if err != nil {
			loggerx.ReqLog.Error(err.Error(), zap.Any("claims", t.Claims.(jwt.MapClaims)))
			ret.SetCodeErr(err)
			c.JSON(http.StatusOK, ret)
			c.Abort()
			return
		}
		if epTime.Time.Before(time.Now()) {
			err := errorx.NewCodeError(module.TokenExpired)
			ret.SetCodeErr(err)
			c.JSON(http.StatusOK, ret)
			c.Abort()
			return
		}
		uid, ok := t.Claims.(jwt.MapClaims)[constx.TokenUID]
		if !ok || uid.(string) == "" {
			loggerx.ReqLog.Error("claim err", zap.Any("claims", t.Claims.(jwt.MapClaims)))
			err := errorx.NewCodeError(module.TokenInvalid)
			ret.SetCodeErr(err)
			c.JSON(http.StatusOK, ret)
			c.Abort()
			return
		}
		tid, ok := t.Claims.(jwt.MapClaims)[constx.TokenTID]
		if !ok || tid.(string) == "" {
			loggerx.ReqLog.Error("claim err", zap.Any("claims", t.Claims.(jwt.MapClaims)))
			err := errorx.NewCodeError(module.TokenInvalid)
			ret.SetCodeErr(err)
			c.JSON(http.StatusOK, ret)
			c.Abort()
			return
		}
		c.Set(constx.TokenUID, uid)
		c.Set(constx.TokenTID, tid)
		c.Next()
	}
}
